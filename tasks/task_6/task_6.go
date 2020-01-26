package main

import (
	"flag"
)

func main() {
	pathFrom := flag.String("from", "", "source file path")
	pathTo := flag.String("to", "", "destination file path")
	limit := flag.Int64("limit", 0, "file data limit")
	offset := flag.Int64("offset", 0, "file data offset")

	flag.Parse()

	//CreateNewFile("1.bin", 2)

	Copy(*pathFrom, *pathTo, *limit, *offset)

	//Copy("1.bin", "2.bin", 300080, 1024*1024*2)
}
