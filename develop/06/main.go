package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

/*

first: separator
second: fields
third: separated only

*/

func main() {
	var (
		fields        int
		delimiter     string
		separatedOnly bool
	)

	flag.IntVar(&fields, "f", 0, "Print only provided field. One space is a default separator.")
	flag.StringVar(&delimiter, "d", "\t", "Change default delimiter.")
	flag.BoolVar(&separatedOnly, "s", false, "Print only lines with separator.")
	flag.Parse()

	filepath := flag.Arg(0)
	fileDsr, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}(fileDsr)

	content, err := io.ReadAll(fileDsr)
	if err != nil {
		log.Fatal(err)
	}

	if len(delimiter) > 1 {
		fmt.Println("Delimiter has to be a single symbol")
		return
	}
	if fields == 0 {
		fmt.Println("You have to specify number of fields")
		return
	}

	lines := strings.Split(string(content), "\n")
	handle(lines, delimiter, fields, separatedOnly)
}

func handle(lines []string, delimiter string, fields int, separated bool) {
	result := make([][]string, 0, cap(lines))
	for _, line := range lines {
		lineFields := strings.Split(line, delimiter)
		if !separated {
			result = append(result, lineFields)
		} else if separated && len(lineFields) > 1 {
			result = append(result, lineFields)
		}
	}

	for j, line := range result {
		for i, field := range line {
			if i == fields-1 {
				fmt.Printf("%s ", field)
			}
		}
		if j < len(result) {
			fmt.Println()
		}
	}
}
