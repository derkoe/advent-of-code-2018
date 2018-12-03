package input

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

// Patch is the elves' patch
type Patch struct {
	id    int
	posX  int
	posY  int
	sizeX int
	sizeY int
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
			id:    atoi(parsedLine[1]),
			posX:  atoi(parsedLine[2]),
			posY:  atoi(parsedLine[3]),
			sizeX: atoi(parsedLine[4]),
			sizeY: atoi(parsedLine[5]),
		}
		res = append(res, patch)
	}
	return res
}

func atoi(str string) int {
	res, _ := strconv.Atoi(str)
	return res
}
