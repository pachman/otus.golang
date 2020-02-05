package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	path := os.Args[1]
	commands := os.Args[2:]

	env, err := ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for key, value := range env {
		fmt.Println(key, value)
	}

	RunCmd(commands, env)
}
