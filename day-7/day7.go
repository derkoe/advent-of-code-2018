package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	one()
}

func one() {
	rules := readRules()
	var result []rune
	for len(rules) > 0 {
		var nextCandidates []rune
		for k, v := range rules {
			if len(v) == 0 {
				nextCandidates = append(nextCandidates, k)
			}
		}
		sort.Slice(nextCandidates, func(i int, j int) bool { return nextCandidates[i] < nextCandidates[j] })

		next := nextCandidates[0]
		result = append(result, next)
		delete(rules, next)
		for k, v := range rules {
			rules[k] = remove(v, next)
		}
	}
	fmt.Println(string(result))
}

func readRules() map[rune][]rune {
	file, _ := os.Open("input")
	scanner := bufio.NewScanner(file)
	rules := make(map[rune][]rune)
	for scanner.Scan() {
		line := scanner.Text()
		step := rune(line[36])
		afterstep := rune(line[5])
		rules[step] = append(rules[step], afterstep)
		if _, found := rules[afterstep]; !found {
			rules[afterstep] = []rune{}
		}
	}
	return rules
}

func remove(s []rune, e rune) []rune {
	for i, a := range s {
		if a == e {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
