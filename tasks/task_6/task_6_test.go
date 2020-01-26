package main

import (
	"log"
	"os"
	"testing"
)

func TestCopyFileSuccess(t *testing.T) {
	source := "test.bin"
	copyFile := "copy-test.bin"

	CreateNewFile(source, 2)

	Copy(source, copyFile, 0, 0)

	fileFrom, _ := os.Open(source)
	fileTo, _ := os.Open(copyFile)

	defer fileTo.Close()
	defer fileFrom.Close()

	statFrom, _ := fileFrom.Stat()
	statTo, _ := fileTo.Stat()

	sizeFrom := statFrom.Size()
	sizeTo := statTo.Size()
	if sizeFrom != sizeTo {
		t.Errorf("Different file size. sizeFrom=%v sizeTo=%v", sizeFrom, sizeTo)
	}
}

func TestCopyLimitSuccess(t *testing.T) {
	source := "test.bin"
	copyFile := "copy-test.bin"

	CreateNewFile(source, 2)

	expected := MbToByte(1)
	Copy(source, copyFile, expected, 0)

	fileFrom, _ := os.Open(source)
	fileTo, _ := os.Open(copyFile)

	defer fileTo.Close()
	defer fileFrom.Close()

	statTo, _ := fileTo.Stat()

	actual := statTo.Size()
	if expected != actual {
		t.Errorf("Different file size actual=%v, expected=%v", actual, expected)
	}
}

func TestCopyCustomLimitSuccess(t *testing.T) {
	source := "test.bin"
	copyFile := "copy-test.bin"

	CreateNewFile(source, 2)

	expected := int64(3000)
	Copy(source, copyFile, expected, 0)

	fileFrom, _ := os.Open(source)
	fileTo, _ := os.Open(copyFile)

	defer fileTo.Close()
	defer fileFrom.Close()

	statTo, _ := fileTo.Stat()

	actual := statTo.Size()
	if expected != actual {
		t.Errorf("Different file size actual=%v, expected=%v", actual, expected)
	}
}

func TestCopyOverOffsetSuccess(t *testing.T) {
	source := "test.bin"
	copyFile := "copy-test.bin"

	CreateNewFile(source, 2)

	offset := MbToByte(2)
	Copy(source, copyFile, 0, offset)

	fileTo, _ := os.Open(copyFile)

	defer fileTo.Close()

	statTo, _ := fileTo.Stat()

	actual := statTo.Size()
	if 0 != actual {
		t.Errorf("Different file size actual=%v", actual)
	}
}

func TestCopyOffsetSuccess(t *testing.T) {
	source := "test.bin"
	copyFile := "copy-test.bin"

	CreateNewFile(source, 2)

	expected := int64(10)
	offset := MbToByte(2) - expected
	Copy(source, copyFile, 0, offset)

	fileTo, _ := os.Open(copyFile)

	defer fileTo.Close()

	statTo, _ := fileTo.Stat()

	actual := statTo.Size()
	if expected != actual {
		t.Errorf("Different file size actual=%v, expected=%v", actual, expected)
	}
}

func MbToByte(megabytes int) int64 {
	return int64(1024 * 1024 * megabytes)
}

func CreateNewFile(path string, megabyteCount int) {
	size := int(MbToByte(megabyteCount))
	b := make([]byte, size)

	for i := 0; i < size; i++ {
		b[i] = 5 //data
	}
	file, _ := os.Create(path)
	defer file.Close()

	_, err := file.Write(b)
	if err != nil {
		log.Panicf("failed to write: %v", err)
	}
}
