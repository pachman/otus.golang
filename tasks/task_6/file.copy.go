package main

import (
	"io"
	"log"
	"os"
)

func Copy(from string, to string, limit int64, offset int64) {
	fileFrom, err := os.Open(from)
	if err != nil {
		panic("Can't open source file " + err.Error())
	}
	defer fileFrom.Close()

	fileTo, err := os.Create(to)
	if err != nil {
		panic("Can't open destination file " + err.Error())
	}
	defer fileTo.Close()

	stat, err := fileFrom.Stat()
	if err != nil {
		panic("Can't read size of source file " + err.Error())
	}

	size := stat.Size()

	if size <= offset {
		return
	}

	if limit <= 0 {
		limit = size
	}

	bufferSize := int64(1024)
	buffer := make([]byte, bufferSize)
	offsetCurrent := offset
	writtenBytes := int64(0)
	for offsetCurrent < offset+limit {
		read, err := fileFrom.ReadAt(buffer, offsetCurrent)
		readBytes := int64(read)
		offsetCurrent += readBytes
		writtenBytes += readBytes

		if writtenBytes > limit {
			writeSize := readBytes - (writtenBytes - limit)
			fileTo.Write(buffer[:writeSize])
		} else if readBytes < bufferSize {
			fileTo.Write(buffer[:readBytes])
		} else {
			fileTo.Write(buffer)
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Panicf("failed to read: %v", err)
		}
	}
}
