package main

import "fmt"

type TaskResult struct {
	Id       int
	Error    error
	TryCount int
}

func (result TaskResult) HasError() bool {
	return result.Error != nil
}

func (result TaskResult) String() string {
	withErrorText := "without errors"
	if result.HasError() {
		withErrorText = "with errors"
	}
	return fmt.Sprintf("Task_%d\tdone %v,\tafter %d attempts", result.Id, withErrorText, result.TryCount)
}

func Run(tasks []func() error, tasksCount int, maxErrorCount int) []TaskResult {
	length := len(tasks)

	if tasksCount > length {
		tasksCount = length
	}

	if maxErrorCount < 0 {
		maxErrorCount = 1
	}

	i := 0

	ch := make(chan TaskResult, length)

	for ; i < tasksCount; i++ {
		index := i
		ch = start(ch, tasks, index, maxErrorCount)
	}

	j := 0
	results := make([]TaskResult, length)

	for result := range ch {
		if length > i {
			ch = start(ch, tasks, i, maxErrorCount)
			i++
		}

		results[j] = result
		j++

		if length == j {
			close(ch)
		}
	}

	return results
}

func start(ch chan TaskResult, tasks []func() error, index int, maxErrorCount int) chan TaskResult {
	go func() {
		tryLeft := maxErrorCount - 1
		for {
			err := tasks[index]()

			if err == nil || tryLeft == 0 {
				ch <- TaskResult{Id: index, Error: err, TryCount: maxErrorCount - tryLeft}
				break
			}

			tryLeft--
		}
	}()

	return ch
}
