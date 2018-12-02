package input

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func Read() []int {
	file, _ := os.Open("input")
	scanner := bufio.NewScanner(file)
	var res []int
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
			os.Exit(2)
		}
		res = append(res, i)
	}
	return res
}
