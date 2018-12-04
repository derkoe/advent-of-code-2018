package main

import (
	"fmt"

	"github.com/derkoe/advent-of-code-2018/day-3"
)

func main() {
	patches := input.Read()
	for _, patch1 := range patches {
		overlaps := false
		for _, patch2 := range patches {
			if patch1.Id != patch2.Id {
				if (patch1.PosX <= patch2.PosX) && (patch1.PosX+patch1.SizeX >= patch2.PosX) ||
					(patch2.PosX <= patch1.PosX) && (patch2.PosX+patch2.SizeX >= patch1.PosX) {
					// x do overlap, check y
					if (patch1.PosY <= patch2.PosY) && (patch1.PosY+patch1.SizeY >= patch2.PosY) ||
						(patch2.PosY <= patch1.PosY) && (patch2.PosY+patch2.SizeY >= patch1.PosY) {
						overlaps = true
					}
				}
			}
		}
		if overlaps == false {
			fmt.Println(patch1)
			return
		}
	}
	fmt.Println("None")
}
