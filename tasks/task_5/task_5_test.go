package main

import (
	"fmt"
	"testing"
)

func TestRunSuccess(t *testing.T) {
	count := 3
	tasks := make([]func() error, count)

	for i := 0; i < count; i++ {
		tasks[i] = func() error {
			return nil
		}
	}

	results := Run(tasks, 1, 1)

	actual := len(results)
	if actual != count {
		t.Errorf("actual=%v, expected=%v", actual, count)
	}

	for i := 0; i < count; i++ {
		if results[i].Error != nil {
			t.Errorf("actual=%v, expected=%v", results[i].HasError(), false)
		}
	}
}

func TestRunMaxErrorCount(t *testing.T) {
	count := 3
	maxErrorCount := 3
	tasks := make([]func() error, count)

	for i := 0; i < count; i++ {
		tasks[i] = func() error {
			return fmt.Errorf("Error ")
		}
	}

	results := Run(tasks, 1, maxErrorCount)

	actual := len(results)
	if actual != count {
		t.Errorf("actual=%v, expected=%v", actual, count)
	}

	for i := 0; i < count; i++ {
		if results[i].Error == nil {
			t.Errorf("actual=%v, expected=%v", results[i].HasError(), false)
		}

		if results[i].TryCount != maxErrorCount {
			t.Errorf("actual=%v, expected=%v", results[i].TryCount, maxErrorCount)
		}
	}
}
