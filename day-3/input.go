package input

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

// Patch is the elves' patch
type Patch struct {
	Id    int
	PosX  int
	PosY  int
	SizeX int
	SizeY int
}

// Read reads input for day 3
func Read() []Patch {
	lineRegexp := regexp.MustCompile("#([0-9]+) @ ([0-9]+),([0-9]+): ([0-9]+)x([0-9]+)")
	file, _ := os.Open("input")
	scanner := bufio.NewScanner(file)
	var res []Patch
	for scanner.Scan() {
		parsedLine := lineRegexp.FindStringSubmatch(scanner.Text())
		patch := Patch{
			Id:    atoi(parsedLine[1]),
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
