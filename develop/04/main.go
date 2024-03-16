package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "пипидастр", "рим", "мир", "мри", "ирм", "рми"}

	anagrams := findAnagram(&words)

	for k := range *anagrams {
		fmt.Println(k, *(*anagrams)[k])
	}
}

func findAnagram(dictionary *[]string) *map[string]*[]string {
	result := make(map[string]*[]string)

	for _, word := range *dictionary {
		lowerWord := strings.ToLower(word)
		result[sortWord([]rune(lowerWord))] = &[]string{}
	}

	for k := range result {
		for _, word := range *dictionary {
			if k == sortWord([]rune(word)) {
				*result[k] = append(*result[k], word)
			}
		}
	}

	clearMap(result)

	m := make(map[string]*[]string, len(result))
	for k := range result {
		m[(*result[k])[0]] = result[k]
	}

	return &m
}

func sortWord(w []rune) string {
	wCopy := make([]rune, len(w))
	sort.Slice(wCopy, func(i, j int) bool {
		return wCopy[i] < wCopy[j]
	})

	return string(wCopy)
}

func clearMap(m map[string]*[]string) {
	for k := range m {
		if len(*m[k]) == 1 {
			delete(m, k)
		}
	}
}
