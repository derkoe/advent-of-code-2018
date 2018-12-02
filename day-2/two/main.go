package main

import (
	"aoc2018/day-2"
	"fmt"
)

func main() {
	input := input.Read()
	for i, s1 := range input {
		for j := i; j < len(input); j++ {
			s2 := input[j]
			if hamming(s1, s2) == 1 {
				printSame(s1, s2)
				return
			}
		}
	}
}

func hamming(s1 string, s2 string) int {
	distance := 0
	r2 := []rune(s2)
	for i, v := range []rune(s1) {
		if r2[i] != v {
			distance++
		}
	}
	return distance
}

func printSame(s1 string, s2 string) {
	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			fmt.Printf("%c", s1[i])
		}
	}
}
