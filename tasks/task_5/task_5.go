package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	count := 30
	tasks := make([]func() error, count)

	for i := 0; i < count; i++ {
		tasks[i] = func() error {
			time.Sleep(2 * time.Second)

			rnd := rand.Int() % 10

			if rnd > 5 {
				return fmt.Errorf("Error!!!")
			}

			return nil
		}
	}

	Run(tasks, 5, 1)
}
