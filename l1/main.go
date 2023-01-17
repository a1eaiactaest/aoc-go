package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func read_file(path string) (lines []string) {
	f, err := os.Open(path)
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	const MAX_CAPACITY = 10000000
	buf := make([]byte, MAX_CAPACITY)
	scanner.Buffer(buf, MAX_CAPACITY)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	check(err)

	return lines
}

func part1(lines []string) int {
	acc := 0
	mcal := 0

	for _, line := range lines {
		acc += strconv.Atoi(line)
	}

	return mcal
}

func main() {
	lines := read_file("input.txt")
	p1 := part1(lines)
	fmt.Println(p1)
}
