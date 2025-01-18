package arrayAndString

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

func MergeAlternately(w, w2 string) string {
	wLength := len(w)
	w2Length := len(w2)
	result := make([]byte, 0, wLength+w2Length)
	for i := 0; i < len(w); i++ {
		if wLength > i {
			result = append(result, w[i])
		}

		if w2Length > i {
			result = append(result, w2[i])
		}
	}

	if w2Length > wLength {
		result = append(result, w2[wLength:]...)
	}

	return string(result)
}

func RomanToInt(s string) int {
	symbols := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	target := 0
	length := len(s)
	for i := 0; length > i; i++ {
		if length == (i + 1) {
			target += symbols[s[i]]
			continue
		}

		if symbols[s[i]] < symbols[s[i+1]] {
			target += -symbols[s[i]]
		} else {
			target += symbols[s[i]]
		}
	}

	return target
}
