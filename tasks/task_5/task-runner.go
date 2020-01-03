package main

func Run(tasks []func() error, tasksCount int, maxErrorCount int) {
	length := len(tasks)

	if tasksCount > length {
		tasksCount = length
	}

	i := 0

	ch := make(chan bool, length)

	for ; i < tasksCount; i++ {
		index := i
		ch = start(tasks, index, ch)
	}

	j := 0
	for isError := range ch {
		if !isError && length > i {
			ch = start(tasks, i, ch)
			i++
		}

		if isError {
			maxErrorCount--
		}

		j++

		if length == j || maxErrorCount == 0 {
			close(ch)
		}
	}
}

func start(tasks []func() error, index int, ch chan bool) chan bool {
	go func() {
		err := tasks[index]()
		ch <- err != nil
	}()

	return ch
}
