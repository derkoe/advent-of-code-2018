package main

import (
	"fmt"

	"github.com/derkoe/advent-of-code-2018/day-1"
)

func main() {
	result := 0
	results := map[int]bool{0: true}
	for {
		for _, i := range input.Read() {
			result += i
			_, found := results[result]
			if found {
				fmt.Println(result)
				return
			}
			results[result] = true
		}
	}
}
