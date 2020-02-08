package main

import (
	"fmt"
	"sync"
)

type TaskResult struct {
	Id    int
	Error error
}

func (result TaskResult) HasError() bool {
	return result.Error != nil
}

func (result TaskResult) String() string {
	withErrorText := "without errors"
	if result.HasError() {
		withErrorText = "with error"
	}
	return fmt.Sprintf("Task_%d\tdone %v", result.Id, withErrorText)
}

func Run(tasks []func() error, goroutinesCount int, maxErrorCount int) []TaskResult {
	tasksCount := len(tasks)
	var wg sync.WaitGroup

	if goroutinesCount > tasksCount {
		goroutinesCount = tasksCount
	}
	if maxErrorCount < 0 {
		maxErrorCount = 1
	}
	ch := make(chan TaskResult, tasksCount)

	taskIndex := 0
	for ; taskIndex < goroutinesCount; taskIndex++ {
		wg.Add(1)
		start(ch, tasks, taskIndex, &wg)
	}

	results := make([]TaskResult, 0, tasksCount)
	errorCount := 0

	for result := range ch {
		results = append(results, result)

		if result.Error != nil {
			errorCount++
		}
		if (tasksCount > taskIndex) && (errorCount < maxErrorCount) {
			wg.Add(1)
			start(ch, tasks, taskIndex, &wg)
			taskIndex++
		} else {
			wg.Wait()
			//close(ch)
		}
	}

	return results
}

func start(ch chan TaskResult, tasks []func() error, taskIndex int, wg *sync.WaitGroup) chan TaskResult {
	go func() {
		defer wg.Done()
		err := tasks[taskIndex]()
		ch <- TaskResult{Id: taskIndex, Error: err}
	}()
	return ch
}
