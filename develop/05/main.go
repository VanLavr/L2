package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strings"
)

/*

stages:
0) number of matches                                                         +
1) add numbers to lines                                                      +
2) ignore case                                                               +
3) fixed                                                                     +
4) inverted (OR below, cannot stack inverted and before or after or context) +
5) after                                                                     +
6) before
7) context

*/

type withLineNumbers struct {
	LineNumber     int
	Line           string
	IgnoreCaseLine string
}

func main() {
	var (
		after      int
		before     int
		context    int
		count      bool
		ignoreCase bool
		invert     bool
		fixed      bool
		lineNum    bool
	)

	flag.IntVar(&after, "A", 0, "print N strings after match")                   //
	flag.IntVar(&before, "B", 0, "print N strings before match")                 //
	flag.IntVar(&context, "C", 0, "print -+N strings around match")              //
	flag.BoolVar(&count, "c", false, "print number of lines within the matches") // +
	flag.BoolVar(&ignoreCase, "i", false, "ignore case")                         // +
	flag.BoolVar(&invert, "V", false, "find all the lines without matches")      // +
	flag.BoolVar(&fixed, "F", false, "find exactly this pattern")                // +
	flag.BoolVar(&lineNum, "n", false, "print line number")                      // +
	flag.Parse()

	substring := flag.Arg(0)
	filepath := flag.Arg(1)

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

	var markedLineNumbers []int
	lines := strings.Split(string(content), "\n")

	for i, line := range lines {
		if strings.Contains(line, substring) {
			markedLineNumbers = append(markedLineNumbers, i+1)
		}
	}

	if count {
		fmt.Println(len(markedLineNumbers))
		return
	}

	addNumbers(lines, lineNum, fixed, invert, ignoreCase, after, before, context, substring)
}

func addNumbers(lines []string, lineNumbers, fixed, invert, ignorecase bool, after, before, context int, substring string) {
	enumerated := make([]withLineNumbers, 0, len(lines))
	if !lineNumbers {
		for _, line := range lines {
			enumerated = append(enumerated, withLineNumbers{LineNumber: -1, Line: line})
		}

		ignoreCase(enumerated, fixed, invert, ignorecase, after, before, context, substring)
		return
	}

	for i, line := range lines {
		enumerated = append(enumerated, withLineNumbers{LineNumber: i + 1, Line: line})
	}

	ignoreCase(enumerated, fixed, invert, ignorecase, after, before, context, substring)
}

func ignoreCase(lines []withLineNumbers, fix, invert, ignorecase bool, after, before, context int, substring string) {
	for i := range lines {
		lines[i].IgnoreCaseLine = strings.ToLower(lines[i].Line)
	}

	fixed(lines, fix, ignorecase, invert, after, before, context, substring)
}

func fixed(lines []withLineNumbers, fix, ignorecase, invert bool, after, before, context int, substring string) {
	if !fix {
		var matches []string
		if !ignorecase {
			for _, line := range lines {
				if strings.Contains(line.Line, substring) {
					matches = append(matches, line.Line)
				}
			}
		} else {
			for _, line := range lines {
				if strings.Contains(line.IgnoreCaseLine, strings.ToLower(substring)) {
					matches = append(matches, line.Line)
				}
			}
		}
		fmt.Println("matches:", matches)

		if invert {
			inverted(lines, matches)
		} else if after > 0 {
			afterLine(lines, matches, after)
		} else if before > 0 {
			beforeLine(lines, matches, before)
		} else if context > 0 {
			aroundLine(lines, matches, context)
		} else {
			printLine(lines, matches)
		}
		return
	}

	var matches []string
	if !ignorecase {
		for _, line := range lines {
			if line.Line == substring {
				matches = append(matches, line.Line)
			}
		}
	} else {
		for _, line := range lines {
			if line.Line == substring || line.Line == strings.ToLower(substring) || line.Line == strings.ToUpper(substring) {
				matches = append(matches, line.Line)
			}
		}
	}
	fmt.Println("matches:", matches)

	if invert {
		inverted(lines, matches)
	} else if after > 0 {
		afterLine(lines, matches, after)
	} else if before > 0 {
		beforeLine(lines, matches, before)
	} else if context > 0 {
		aroundLine(lines, matches, context)
	} else {
		printLine(lines, matches)
	}
}

func inverted(lines []withLineNumbers, matches []string) {
	for _, match := range matches {
		for i := len(lines) - 1; i > -1; i-- {
			if match == lines[i].Line {
				lines = slices.Delete(lines, i, i+1)
			}
		}
	}

	if lines[0].LineNumber < 0 {
		for _, line := range lines {
			fmt.Println(line.Line)
		}
	} else {
		for _, line := range lines {
			fmt.Printf("%d: %s\n", line.LineNumber, line.Line)
		}
	}
}

func print(matches []string) {
	for _, match := range matches {
		fmt.Println(match)
	}
}

func afterLine(lines []withLineNumbers, matches []string, after int) {
	j := 0
	for i := 0; i < len(lines); i++ {
		if lines[i].Line == matches[j] {
			for k := i; k <= i+after; k++ {
				if k > len(lines)-1 {
					break
				}
				if lines[k].LineNumber > 0 {
					fmt.Printf("%d: %s\n", lines[k].LineNumber, lines[k].Line)
				} else {
					fmt.Println(lines[k].Line)
				}
			}
			fmt.Println("...")
			fmt.Println()
			j++
			if j > len(matches)-1 {
				break
			}
		}
	}
}

func beforeLine(lines []withLineNumbers, matches []string, before int) {
	j := 0
	for i := 0; i < len(lines); i++ {
		if lines[i].Line == matches[j] {
			for k := i - before; k <= i; k++ {
				if k < 0 {
					k = 0
				}
				if lines[k].LineNumber > 0 {
					fmt.Printf("%d: %s\n", lines[k].LineNumber, lines[k].Line)
				} else {
					fmt.Println(lines[k].Line)
				}
			}
			fmt.Println("...")
			fmt.Println()
			j++
			if j > len(matches)-1 {
				break
			}
		}
	}
}

func aroundLine(lines []withLineNumbers, matches []string, around int) {
	j := 0
	for i := 0; i < len(lines); i++ {
		if lines[i].Line == matches[j] {
			for k := i - around; k <= i+around; k++ {
				if k < 0 {
					k = 0
				}
				if k > len(lines)-1 {
					break
				}
				if lines[k].LineNumber > 0 {
					fmt.Printf("%d: %s\n", lines[k].LineNumber, lines[k].Line)
				} else {
					fmt.Println(lines[k].Line)
				}
			}
			fmt.Println("...")
			fmt.Println()
			j++
			if j > len(matches)-1 {
				break
			}
		}
	}
}

func printLine(lines []withLineNumbers, matches []string) {
	if lines[0].LineNumber == -1 {
		i := 0
		for _, line := range lines {
			if line.Line == matches[i] {
				fmt.Println(line.Line)
				i++
				if i > len(matches)-1 {
					break
				}
			}
		}
	} else {
		i := 0
		for _, line := range lines {
			if line.Line == matches[i] {
				fmt.Printf("%d: %s\n", line.LineNumber, line.Line)
				i++
				if i > len(matches)-1 {
					break
				}
			}
		}
	}
}
