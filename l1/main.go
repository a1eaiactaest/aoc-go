package main

import "fmt"

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
      cals += atoi(line)
    }
	} 
	return max_cals
}

/*
func part2(lines []string) int {
  return lines
}
*/

func main() {
	lines := file_to_lines("input.txt")
	p1 := part1(lines)
	fmt.Println(p1)
}
