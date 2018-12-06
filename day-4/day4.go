package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"time"
)

// Guard holds data about a sleepy guard
type Guard struct {
	ID        int
	MinuteSum int
	MinuteMap map[int]int
}

func main() {
	guards := readGuards()
	fmt.Println("--- Part One ---")
	one(guards)
	fmt.Println("--- Part Two ---")
	two(guards)
}

func one(guards []Guard) {
	var sleepyGuard Guard
	maxMinutes := 0
	for _, guard := range guards {
		if guard.MinuteSum > maxMinutes {
			maxMinutes = guard.MinuteSum
			sleepyGuard = guard
		}
	}
	printSleepyGuard(sleepyGuard)
}

func two(guards []Guard) {
	maxCount := 0
	var sleepyGuard Guard
	for _, guard := range guards {
		for _, count := range guard.MinuteMap {
			if count > maxCount {
				maxCount = count
				sleepyGuard = guard
			}
		}
	}
	printSleepyGuard(sleepyGuard)
}

func printSleepyGuard(sleepyGuard Guard) {
	maxCount := 0
	var maxMin int
	for min, count := range sleepyGuard.MinuteMap {
		if count > maxCount {
			maxMin = min
			maxCount = count
		}
	}
	fmt.Printf("Most sleepy guard %v - most asleep in minute %v => %v\n", sleepyGuard.ID, maxMin, sleepyGuard.ID*maxMin)
}

func readGuards() []Guard {
	file, _ := os.Open("input")
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	sort.Strings(lines)
	lineRegexp := regexp.MustCompile("#([0-9]+)")
	guardMap := make(map[int]Guard)
	var currentGuardID int
	var asleepTime time.Time
	for _, line := range lines {
		timestamp, _ := time.Parse("2006-01-02 15:04", string(line[1:17]))
		switch line[19] {
		case 'G': // Guard #nr begins shift
			currentGuardID = atoi(lineRegexp.FindStringSubmatch(line)[1])
		case 'f': // falls asleep
			asleepTime = timestamp
		case 'w': // wakes up
			var guard Guard
			if g, exists := guardMap[currentGuardID]; exists {
				guard = g
			} else {
				guard = Guard{
					ID:        currentGuardID,
					MinuteSum: 0,
					MinuteMap: make(map[int]int),
				}
			}
			guard.MinuteSum += int(timestamp.Sub(asleepTime).Minutes())

			oneMinute, _ := time.ParseDuration("1m")
			for curMinute := asleepTime; curMinute.Before(timestamp); curMinute = curMinute.Add(oneMinute) {
				guard.MinuteMap[curMinute.Minute()]++
			}
			guardMap[currentGuardID] = guard
		}
	}

	return values(guardMap)
}

func atoi(str string) int {
	res, _ := strconv.Atoi(str)
	return res
}

func values(guardMap map[int]Guard) []Guard {
	var guards []Guard
	for _, guard := range guardMap {
		guards = append(guards, guard)
	}
	return guards
}
