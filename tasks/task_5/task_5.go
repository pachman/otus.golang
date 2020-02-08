package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	count := 125
	tasks := make([]func() error, count)

	for i := 0; i < count; i++ {
		tasks[i] = func() error {
			time.Sleep(300 * time.Millisecond)

			rand.Seed(time.Now().UnixNano())
			rnd := rand.Int() % 10

			fmt.Println(rnd)
			if rnd > 4 {
				return fmt.Errorf("Error %d ", rnd)
			}
			return nil
		}
	}

	results := Run(tasks, 15, 30)

	fmt.Println("List of results:")
	for i := 0; i < len(results); i++ {
		fmt.Println(results[i])
	}
}
