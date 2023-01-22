package algo

// read: https://en.wikipedia.org/wiki/Combination

func CombinationsInts(nums []int, targetLength int) [][]int {
	if targetLength > len(nums) {
		panic("target length is greater than length of input slice")
	}

	var combos [][]int

	for i := 0; i < len(nums); i++ {
		// current is empty
		combos = append(combos, helperCombos(nums[i:], targetLength, []int{})...)
	}

	return combos
}

func helperCombos(nums []int, length int, current []int) [][]int {
	if len(current) == length {
		return [][]int{append([]int{}, current...)}
	}

	var combos [][]int

	for i, v := range nums {
		// append value to current combo
		current = append(current, v)

		recurseRet := helperCombos(nums[i+1:], length, current)
		combos = append(combos, recurseRet...)

		// backtrack
		current = current[:len(current)-1]
	}
	return combos
}