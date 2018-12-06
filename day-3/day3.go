package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// Patch is the elves' patch
type Patch struct {
	ID    int
	PosX  int
	PosY  int
	SizeX int
	SizeY int
}

func main() {
	patches := readPatches()
	fmt.Println("--- Part One ---")
	one(patches)
	fmt.Println("--- Part Two ---")
	two(patches)
}

func one(patches []Patch) {
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

func two(patches []Patch) {
	for _, patch1 := range patches {
		overlaps := false
		for _, patch2 := range patches {
			if patch1.ID != patch2.ID {
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
			fmt.Println(patch1.ID)
			return
		}
	}
	fmt.Println("None")
}

// readPatches reads input for day 3
func readPatches() []Patch {
	lineRegexp := regexp.MustCompile("#([0-9]+) @ ([0-9]+),([0-9]+): ([0-9]+)x([0-9]+)")
	file, _ := os.Open("input")
	scanner := bufio.NewScanner(file)
	var res []Patch
	for scanner.Scan() {
		parsedLine := lineRegexp.FindStringSubmatch(scanner.Text())
		patch := Patch{
			ID:    atoi(parsedLine[1]),
			PosX:  atoi(parsedLine[2]),
			PosY:  atoi(parsedLine[3]),
			SizeX: atoi(parsedLine[4]),
			SizeY: atoi(parsedLine[5]),
		}
		res = append(res, patch)
	}
	return res
}

func atoi(str string) int {
	res, _ := strconv.Atoi(str)
	return res
}
