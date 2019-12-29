package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"unicode"
)

func main() {
	input := `\32a4bc2d5e\\5`
	//input := `\a4`

	unpack, err := Unpack(input)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(unpack)
}

func Unpack(packedString string) (unpackedString string, err error) {

	if len(packedString) == 0 || unicode.IsDigit(rune(packedString[0])) {
		err = errors.New("Invalid input string")
		return
	}

	var buffer strings.Builder
	var unpackedChar rune
	isCharEscaped := false
	for _, char := range packedString {

		if isCharEscaped {
			if isEscape(char) || unicode.IsDigit(char) {
				unpackedChar = char
				buffer.WriteRune(char)
				isCharEscaped = false
				continue
			} else {
				err = errors.New("Invalid input string")
				break
			}
		}

		if isEscape(char) {
			isCharEscaped = true
			continue
		}

		if unicode.IsDigit(char) {
			count := int(char - 49)
			buffer.WriteString(strings.Repeat(string(unpackedChar), count))
		} else {
			unpackedChar = char
			buffer.WriteRune(unpackedChar)
		}
	}
	if err == nil {
		unpackedString = buffer.String()
	} else {
		unpackedString = ""
	}
	return unpackedString, err
}

func isEscape(b rune) bool {
	return b == '\\'
}
