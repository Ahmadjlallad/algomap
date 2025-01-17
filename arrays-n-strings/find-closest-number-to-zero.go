package arraysAndStrings

import (
	"math"
)

func FindClosestNumber(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	target := nums[0]
	for _, num := range nums {
		currentNum := int(math.Abs(float64(num)))
		absTarget := int(math.Abs(float64(target)))
		if absTarget == num {
			target = currentNum
		} else if currentNum < absTarget {
			target = num
		}
	}

	return target
}
