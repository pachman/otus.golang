package main

import (
	"io"
	"log"
	"os"
)

func main() {
	CreateNewFile("1.bin")

	Copy("1.bin", "2.bin", 300080, 1024)
}

func Copy(from string, to string, limit int64, offset int64) {
	fileFrom, _ := os.Open(from) //todo error
	fileTo, _ := os.Create(to)   //todo error

	N := int64(1024)
	buf := make([]byte, N)
	offsetCurrent := offset
	for offsetCurrent < offset+limit {
		read, err := fileFrom.ReadAt(buf, offsetCurrent)
		offsetCurrent += int64(read)

		fileTo.Write(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Panicf("failed to read: %v", err)
		}
	}

	fileTo.Close()
	fileFrom.Close()
}

func CreateNewFile(path string) {
	size := 1024 * 1024 * 10
	b := make([]byte, size) // заполнен нулями

	for i := 0; i < size; i++ {
		b[i] = 5
	}
	file, _ := os.Create(path)
	_, err := file.Write(b)
	if err != nil {
		log.Panicf("failed to write: %v", err)
	}

	// мы записали 1M данных !
	file.Close()
}
