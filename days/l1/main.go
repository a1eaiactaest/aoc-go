package main

import (
	"fmt"
	"sort"

	"github.com/a1eaiactaest/aoc-go/cast"
	"github.com/a1eaiactaest/aoc-go/files"
	"github.com/a1eaiactaest/aoc-go/maths"
)

func part1(lines []string) int {
  cals := 0
	max_cals := 0

	for _, line := range lines {
    if line == "" {
      if cals > max_cals { // comp
        max_cals = cals 
      }
      cals = 0
    } else { // acc
      cals += cast.ToInt(line)
    }
	} 
	return max_cals
}

func part2(lines []string) int {
	/*
	Idea:
		* accumulate elf's cals
		* sort
		* pick last three
	*/

	cals := []int{}
	thisElf := 0
	for _, line := range lines {
		if line == "" {
			cals = append(cals, thisElf)
			thisElf = 0
		} else {
			thisElf += cast.ToInt(line)
		}
	}

	sort.Ints(cals)
	topThreeSum := maths.SumIntSlice(cals[len(cals)-3:])
	return topThreeSum
}

func main() {
	lines := files.FileToLines("input.txt")
	p1 := part1(lines)
	p2 := part2(lines)
	fmt.Println(p1)
	fmt.Println(p2)
}
