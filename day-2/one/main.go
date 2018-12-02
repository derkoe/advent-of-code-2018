package main

import (
	"aoc2018/day-2"
	"fmt"
)

func main() {
	contains2Count := 0
	contains3Count := 0
	for _, s := range input.Read() {
		charCount := charCount(s)
		if contains(charCount, 2) {
			contains2Count++
		}
		if contains(charCount, 3) {
			contains3Count++
		}
	}
	fmt.Println(contains2Count * contains3Count)
}

func contains(charCount map[int32]int, i int) bool {
	for _, v := range charCount {
		if v == i {
			return true
		}
	}
	return false
}

func charCount(s string) map[int32]int {
	results := make(map[int32]int)
	for _, c := range s {
		results[c] = results[c] + 1
	}
	return results
}
