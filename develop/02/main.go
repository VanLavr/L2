package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"unicode"
)

var errInvalidString = errors.New("provided string is invalid")

func main() {
	res, err := unpackString("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}

func unpackString(s string) (string, error) {
	var result string

	runes := []rune(s)
	if len(runes) == 0 {
		return "", nil
	}

	if unicode.IsDigit(runes[0]) {
		return "", errInvalidString
	}

	for i := 0; i < len(runes)-1; i++ {
		if unicode.IsDigit(runes[i+1]) {
			len, err := strconv.Atoi(string(runes[i+1]))
			if err != nil {
				log.Fatal("something went wrong")
			}
			for j := 0; j < len; j++ {
				result += string(runes[i])
			}

		} else if !unicode.IsDigit(runes[i]) {
			result += string(runes[i])
		}
	}

	if !unicode.IsDigit(runes[len(runes)-1]) {
		result += string(runes[len(runes)-1])
	}

	return result, nil
}
