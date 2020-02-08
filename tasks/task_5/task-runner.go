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

	closeChannel := make(chan bool, 1)
	taskChannel := make(chan TaskResult, goroutinesCount)

	taskIndex := 0
	for ; taskIndex < goroutinesCount; taskIndex++ {
		start(taskChannel, tasks, taskIndex, closeChannel, wg)
	}

	results := make([]TaskResult, 0, tasksCount)
	errorCount := 0

	for result := range taskChannel {
		results = append(results, result)

		if result.Error != nil {
			errorCount++

			if errorCount >= maxErrorCount {
				closeChannel <- true

				break
			}
		}

		if tasksCount > taskIndex {
			start(taskChannel, tasks, taskIndex, closeChannel, wg)
			taskIndex++
		}

		if tasksCount == len(results) {
			break
		}

	}

	//close(closeChannel)

	fmt.Println("finish")

	close(taskChannel)
	close(closeChannel)
	return results
}

func start(ch chan TaskResult, tasks []func() error, taskIndex int, closeChannel chan bool, wg sync.WaitGroup) {
	wg.Add(1)
	//closeChannel <- false
	go func() {
		defer wg.Done()
		err := tasks[taskIndex]()

		ch <- TaskResult{Id: taskIndex, Error: err}
	}()
}
