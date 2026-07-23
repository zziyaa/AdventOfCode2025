package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/zziyaa/aoc25/util"
)

func main() {
	input := util.ReadLines(os.Stdin)[0]
	part1, part2 := solve(input)
	fmt.Printf("Part 1: %d\nPart 2: %d\n", part1, part2)
}

func solve(in string) (part1 int, part2 int) {
	for part := range strings.SplitSeq(in, ",") {
		r := strings.Split(part, "-")
		start, _ := strconv.Atoi(r[0])
		end, _ := strconv.Atoi(r[1])

		for id := start; id <= end; id++ {
			if !isValidPart1(id) {
				part1 += id
			}

			if !isValidPart2(id) {
				part2 += id
			}
		}
	}
	return
}

func isValidPart1(productID int) bool {
	id := strconv.Itoa(productID)
	return !isRepeating(id, 2)
}

func isValidPart2(productID int) bool {
	id := strconv.Itoa(productID)
	for i := 2; i <= len(id); i++ {
		if isRepeating(id, i) {
			return false
		}
	}
	return true
}

func isRepeating(s string, partCount int) bool {
	if len(s)%partCount != 0 {
		return false
	}

	partSize := len(s) / partCount
	for i := partSize; i < len(s); i += partSize {
		if s[i:i+partSize] != s[:partSize] {
			return false
		}
	}
	return true
}
