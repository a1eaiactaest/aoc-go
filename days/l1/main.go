package main

import (
	"fmt"

	"github.com/a1eaiactaest/aoc-go/cast"
	"github.com/a1eaiactaest/aoc-go/files"
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

func main() {
	lines := files.FileToLines("input.txt")
	p1 := part1(lines)
	fmt.Println(p1)
}
