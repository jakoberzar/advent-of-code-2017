package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	buf, _ := ioutil.ReadFile("input.txt")
	s := string(buf)
	lines := strings.Split(strings.TrimSpace(s), "\n")

	part1 := 0
	part2 := 0
	for _, line := range lines {
		lineNumbers := strings.Split(strings.TrimSpace(line), "\t")

		part1 += maxMinusMin(lineNumbers)
		part2 += dividables(lineNumbers)
	}

	println(part1)
	println(part2)
}

func maxMinusMin(lineNumbers []string) int {
	max, _ := strconv.Atoi(lineNumbers[0])
	min, _ := strconv.Atoi(lineNumbers[0])
	for _, number := range lineNumbers {
		numberInt, _ := strconv.Atoi(number)
		if numberInt > max {
			max = numberInt
		}
		if numberInt < min {
			min = numberInt
		}
	}
	return max - min
}

func dividables(lineNumbers []string) int {
	for _, number := range lineNumbers {
		numberInt, _ := strconv.Atoi(number)
		for _, number2 := range lineNumbers {
			numberInt2, _ := strconv.Atoi(number2)
			if numberInt%numberInt2 == 0 && numberInt != numberInt2 {
				return numberInt / numberInt2
			}
		}

	}
	return 0
}
