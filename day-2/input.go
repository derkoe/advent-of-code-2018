package input

import (
	"bufio"
	"os"
)

// Reads input for day 2
func Read() []string {
	file, _ := os.Open("input")
	scanner := bufio.NewScanner(file)
	var res []string
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	return res
}
