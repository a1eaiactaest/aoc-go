package maths

import "math"

func MaxInt(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v > ret {
			ret = v
		}
	}
	return ret
}

func MinInt(nums ...int) int {
	ret := nums[0]

	for _, v := range nums {
		if v < ret {
			ret = v
		}
	}
	return ret
}

func AbsoluteInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func SumIntSlice(nums []int) int {
	var acc int
	for _, v := range nums {
		acc += v
	}
	return acc
}

func ProdIntSlice(nums []int) int {
	var ret int
	for _, v := range nums {
		ret *= v
	}
	return ret
}


func PythagoreanDistance(x1, y1, x2, y2 int) float64 {
	x := float64(x1 - x2)
	y := float64(y1 - y2)

	squares := math.Pow(x, 2) + math.Pow(y, 2)
	return math.Sqrt(squares)
}

func ManhattanDistance(x1, y1, x2, y2 int) int {
	x := x1 - x2
	y := y1 - y2

	if x < 0 {
		x *= -1
	}

	if y < 0 {
		y *= -1
	}

	return x + y
}