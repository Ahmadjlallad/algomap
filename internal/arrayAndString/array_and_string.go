package arrayAndString

import (
	"fmt"
	"math"
	"slices"
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

func IsSubsequence(s, t string) bool {
	sLen := len(s)
	tLen := len(t)
	sPointer := 0

	if sLen == 0 {
		return true
	}

	if sLen > tLen {
		return false
	}

	for i := 0; i < tLen; i++ {
		if sPointer == sLen {
			break
		}

		if t[i] == s[sPointer] {
			sPointer++
		}
	}

	return sPointer == sLen
}

/*
	func MaxProfit(prices []int) int {
		length := len(prices)

		if length == 0 {
			return 0
		}

		buy := prices[0]
		sell := 0
		currentMax := 0
		for i, price := range prices {
			if price > buy && price > sell && currentMax <= (price-buy) {
				sell = price
				currentMax = sell - buy
			} else if i != length-1 && price < buy {
				buy = price
				sell = 0
			}
		}

		if sell == 0 && currentMax == 0 {
			return 0
		}

		return currentMax
	}
*/
func MaxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		}

		currentProfit := price - minPrice

		if currentProfit > maxProfit {
			maxProfit = currentProfit
		}
	}

	return maxProfit
}

func LongestCommonPrefix(strs []string) string {
	length := len(strs)
	prefix := ""
	currentPrefix := ""
	arrayIndex := 0
	stringIndex := 0
	for {
		if arrayIndex >= length {
			prefix += currentPrefix
			arrayIndex = 0
			stringIndex++
			currentPrefix = ""
		}

		if len(strs[arrayIndex])-1 < stringIndex {
			break
		}

		if currentPrefix == "" {
			currentPrefix = string(strs[arrayIndex][stringIndex])
		}

		if string(strs[arrayIndex][stringIndex]) != currentPrefix {
			break
		}

		currentPrefix = string(strs[arrayIndex][stringIndex])
		arrayIndex++
	}

	return prefix
}

func SummaryRanges(nums []int) []string {
	length := len(nums)
	if length == 0 {
		return []string{}
	}

	result := make([]string, 0, length)
	start := nums[0]
	previous := nums[0]

	for i := 1; i < length; i++ {
		if previous+1 == nums[i] {
			previous = nums[i]
			continue
		}

		if previous == start {
			result = append(result, fmt.Sprintf("%d", start))
		} else {
			result = append(result, fmt.Sprintf("%d->%d", start, previous))
		}

		start = nums[i]
		previous = nums[i]
	}

	if start == previous {
		result = append(result, fmt.Sprintf("%d", start))
	} else {
		result = append(result, fmt.Sprintf("%d->%d", start, previous))
	}

	return result
}

func ProductExceptSelf(nums []int) []int {
	numsLen := len(nums)
	left := make([]int, numsLen)
	right := make([]int, numsLen)
	multiplier := 1

	for i := 0; i < numsLen; i++ {
		left[i] = multiplier
		multiplier *= nums[i]
	}

	multiplier = 1
	for i := numsLen - 1; i >= 0; i-- {
		right[i] = multiplier
		multiplier *= nums[i]
	}

	for i := 0; i < numsLen; i++ {
		nums[i] = right[i] * left[i]
	}

	return nums
}

// Merge [[1,3],[2,6],[8,10],[15,18]]
// result [[1,6],[8,10],[15,18]]
func Merge(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(a, b []int) int {
		if a[0] >= b[0] {
			return 1
		}

		return -1
	})

	var out [][]int
	start := intervals[0][0]
	end := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		interval := intervals[i]

		if interval[1] >= start && interval[0] <= end {
			if end < interval[1] {
				end = interval[1]
			}

			continue
		}

		out = append(out, []int{start, end})
		start = interval[0]
		end = interval[1]
	}

	out = append(out, []int{start, end})

	return out
}
