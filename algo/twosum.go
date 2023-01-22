package algo

func TwoSum(nums []int, target int) (n1 int, n2 int, found bool) {
	seen := make(map[int]bool, len(nums))

	for _, v := range nums {
		if seen[target-v] {
			return target - v, v, true
		}
		seen[v] = true
	}
	return 0, 0, false
}