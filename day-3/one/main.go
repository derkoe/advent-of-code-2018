package main

import (
	"fmt"

	"github.com/derkoe/advent-of-code-2018/day-3"
)

func main() {
	patches := input.Read()
	fabric := [1000][1000]int{}
	for _, patch := range patches {
		for x := patch.PosX; x < patch.PosX+patch.SizeX; x++ {
			for y := patch.PosY; y < patch.PosY+patch.SizeY; y++ {
				fabric[x][y]++
			}
		}
	}
	count := 0
	for _, line := range fabric {
		for _, square := range line {
			if square > 1 {
				count++
			}
		}
	}
	fmt.Println(count)
}
