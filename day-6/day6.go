package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	x int
	y int
}

func main() {
	coords, maxX, maxY := readCoords()
	fmt.Println("--- Part One ---")
	one(coords, maxX, maxY)
	fmt.Println("--- Part Two ---")
	two(coords, maxX, maxY)
}

func one(coords []Coord, maxX int, maxY int) {
	field := make([][]int, maxY+1)
	count := make(map[int]int)
	ignore := make(map[int]bool)
	for y := 0; y <= maxY; y++ {
		field[y] = make([]int, maxX+1)
		for x := 0; x <= maxX; x++ {
			minDistance := 10000
			closestTo := 0
			for id, c := range coords {
				distance := abs(x-c.x) + abs(y-c.y)
				if distance < minDistance {
					minDistance = distance
					closestTo = id + 1
				} else if distance == minDistance {
					closestTo = 0
				}
			}
			if closestTo > 0 {
				count[closestTo]++
				// edge field -> infinity
				if x == 0 || y == 0 || x == maxX || y == maxY {
					ignore[closestTo] = true
				}
			}
		}
	}
	var maxVal int
	for k, v := range count {
		if _, found := ignore[k]; !found {
			if v > maxVal {
				maxVal = v
			}
		}
	}

	fmt.Println(maxVal)
}

func two(coords []Coord, maxX int, maxY int) {
	regionSize := 0
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			distanceSum := 0
			for _, c := range coords {
				distance := abs(x-c.x) + abs(y-c.y)
				distanceSum += distance
			}
			if distanceSum < 10000 {
				regionSize++
			}
		}
	}
	fmt.Println(regionSize)
}

func readCoords() ([]Coord, int, int) {
	file, _ := os.Open("input")
	scanner := bufio.NewScanner(file)
	var coords []Coord
	var maxX, maxY int
	for scanner.Scan() {
		coordStrings := strings.Split(scanner.Text(), ", ")
		x := atoi(coordStrings[0])
		y := atoi(coordStrings[1])
		coords = append(coords, Coord{
			x: x,
			y: y,
		})
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
	}

	return coords, maxX, maxY
}

func atoi(str string) int {
	res, _ := strconv.Atoi(str)
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
