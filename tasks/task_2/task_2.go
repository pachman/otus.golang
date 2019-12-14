package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
)

func main() {
	input := `a4bc2d5e`

	unpack, err := Unpack(input)

	if err == nil {
		fmt.Println(unpack)
	} else {
		log.Fatal(err)
	}

}

func Unpack(s string) (result string, err error) {
	var buffer bytes.Buffer

	var current byte

	if len(s) == 0 || isNumber(s[0]) {
		err = errors.New("Invalid input string")
		return
	}

	escape := false

	for i := 0; i < len(s); i++ {
		b := s[i]

		if escape && (isEscape(b) || isNumber(b)) {
			current = s[i]
			buffer.WriteByte(current)
			escape = false
			continue
		}

		if escape {
			buffer.WriteByte(b)
			escape = false
			continue
		}

		if isEscape(b) {
			escape = true
			continue
		}

		if isNumber(b) {
			count := int(b - 49)
			for j := 0; j < count; j++ {
				buffer.WriteByte(current)
			}
		} else {
			current = s[i]
			buffer.WriteByte(current)
		}
	}

	result = buffer.String()
	err = nil

	return
}

func isEscape(b uint8) bool {
	return b == 92
}

func isNumber(b uint8) bool {
	return b > 48 && b < 58
}
