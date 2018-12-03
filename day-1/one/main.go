package main

import (
	"fmt"

	"github.com/derkoe/advent-of-code-2018/day-1"
)

func main() {
	result := 0
	for _, i := range input.Read() {
		result += i
	}
	fmt.Println(result)
}
