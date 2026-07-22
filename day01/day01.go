package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/zziyaa/aoc25/util"
)

func main() {
	lines := util.ReadLines(os.Stdin)
	part1, part2 := solve(lines)
	fmt.Printf("Part 1: %d\nPart 2: %d\n", part1, part2)
}

func solve(lines []string) (part1 int, part2 int) {
	pos := 50

	for _, l := range lines {
		d, _ := strconv.Atoi(l[1:])
		for range d {
			if string(l[0]) == "R" {
				pos++
				if pos >= 100 {
					pos = 0
				}
			} else {
				pos--
				if pos < 0 {
					pos = 99
				}
			}

			if pos == 0 {
				part2++
			}
		}

		if pos == 0 {
			part1++
		}
	}

	return
}
