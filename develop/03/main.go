package main

import (
	"cmp"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var (
		uniqueRequired  bool
		reverseRequired bool
		numericRequired bool
		column          int
		filename        string
	)

	flag.BoolVar(&uniqueRequired, "u", false, "prints only unique strings")
	flag.BoolVar(&reverseRequired, "r", false, "prints reverse sorted strings")
	flag.BoolVar(&numericRequired, "n", false, "stands for sorting files with numbers within")
	flag.IntVar(&column, "k", 1, "stands for sorting depending on what column provided")
	flag.Parse()
	filename = flag.Arg(0)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")
	lines = append(lines[:len(lines)-1], lines[len(lines):]...)

	// handle -u
	if uniqueRequired {
		lines = createUnique(lines)
	}

	// handle -n and -k
	if numericRequired {
		sort.Slice(lines, func(i, j int) bool {
			a := strings.Fields(lines[i])[column-1]
			b := strings.Fields(lines[j])[column-1]
			dig1, err1 := strconv.ParseFloat(a, 64)
			dig2, err2 := strconv.ParseFloat(b, 64)
			if err1 == nil && err2 == nil {
				return dig1 < dig2
			}
			return true
		})
		if !reverseRequired {
			for _, line := range lines {
				fmt.Println(line)
			}
			return
		}
		for i := range lines {
			fmt.Println(lines[len(lines)-i-1])
		}
		return
	}

	// handle -k
	slices.SortFunc(lines, func(a, b string) int {
		if reverseRequired {
			if a != "" && b != "" {
				return -1 * cmp.Compare(strings.ToLower(strings.Fields(a)[column-1]), strings.ToLower(strings.Fields(b)[column-1]))
			}
			return -1 * cmp.Compare(strings.ToLower(a), strings.ToLower(b))
		}
		if a != "" && b != "" {
			return cmp.Compare(strings.ToLower(strings.Fields(a)[column-1]), strings.ToLower(strings.Fields(b)[column-1]))
		}
		return cmp.Compare(strings.ToLower(a), strings.ToLower(b))
	})

	for _, line := range lines {
		fmt.Println(line)
	}
}

func createUnique(lines []string) []string {
	result := []string{}
	unique := make(map[string]struct{}, len(lines))

	for _, line := range lines {
		unique[line] = struct{}{}
	}

	for k := range unique {
		result = append(result, k)
	}

	sort.Strings(result)
	return result
}
