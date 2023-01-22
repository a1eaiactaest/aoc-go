package algo

import "strings"

func PermuteIntSlice(numbers []int) [][]int {
	return recurseIntSlice(numbers, 0)
}

// helper for the above
func recurseIntSlice(numbers []int, startIndex int) [][]int {
	if startIndex == len(numbers) {
		return [][]int{append([]int{}, numbers...)}
	}
	var perms [][]int
	for i := startIndex; i < len(numbers); i++ {
		numbers[startIndex], numbers[i] = numbers[i], numbers[startIndex]
		perms = append(perms, recurseIntSlice(numbers, startIndex+1)...)
		numbers[startIndex], numbers[i] = numbers[i], numbers[startIndex]
	}
	return perms
}

func PermuteString(str string) []string {
	return recurseString(strings.Split(str, ""), 0)
}

func recurseString(sli []string, index int) []string {
	if index == len(sli) {
		return []string{strings.Join(sli, "")}
	}

	var perms []string

	for i := index; i < len(sli); i++ {
		sli[i], sli[index] = sli[index], sli[i]
		perms = append(perms, recurseString(sli, index+1)...)
		sli[i], sli[index] = sli[index], sli[i]
	}
	return perms
}

func PermuteStringSlice(in []string) [][]string {
	return recurseStringSlice(in, 0)
}

func recurseStringSlice(in []string, startIndex int) [][]string {
	if startIndex == len(in) {
		// makes a copy
		return [][]string{append([]string{}, in...)}
	}

	var perms [][]string
	for i := startIndex; i < len(in); i++ {
		in[startIndex], in[i] = in[i], in[startIndex]
		perms = append(perms, recurseStringSlice(in, startIndex+1)...)
		in[startIndex], in[i] = in[i], in[startIndex]
	}
	return perms
}