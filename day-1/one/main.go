package main

import (
	"aoc2018/day-1"
	"fmt"
)

func main() {
	result := 0
	for _, i := range input.Read() {
		result += i
	}
	fmt.Println(result)
}
