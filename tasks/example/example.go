package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var msg string
	flag.StringVar(&msg, "msg", "", "message to print")
	flag.Parse()

	//env := os.Environ() // слайс строк "key=value"
	fmt.Println(msg)
	fmt.Println(os.ExpandEnv("$USER lives in ${CITY}"))
}
