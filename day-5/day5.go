package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

const units = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	fmt.Println("--- Part One ---")
	polymer := readPolymer()
	one(polymer)
	fmt.Println("--- Part Two ---")
	two(polymer)
}

func one(polymer string) {
	fmt.Println(react(polymer))
}

func two(polymer string) {
	var shortestPolymer int
	for _, unit := range units {
		unitLower := unicode.ToLower(unit)
		reducedPolymer := strings.Map(func(r rune) rune {
			if r == unit || r == unitLower {
				return -1
			}
			return r
		}, polymer)

		size := react(reducedPolymer)
		if shortestPolymer == 0 || size < shortestPolymer {
			shortestPolymer = size
		}
	}
	fmt.Println(shortestPolymer)
}

func react(polymer string) int {
outer:
	for {
		var prevChar rune
		for i, c := range polymer {
			if c != prevChar && unicode.ToLower(c) == unicode.ToLower(prevChar) {
				if i+1 >= len(polymer) {
					polymer = polymer[0 : i-1]
				} else {
					polymer = polymer[:i-1] + polymer[i+1:]
				}
				continue outer
			}
			prevChar = c
		}
		return len(polymer) - 1
	}
}

func readPolymer() string {
	contents, _ := ioutil.ReadFile("input")
	return string(contents)
}
