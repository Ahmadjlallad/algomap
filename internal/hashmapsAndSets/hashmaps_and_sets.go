package hashmapsAndSets

func NumJewelsInStones(jewels string, stones string) int {
	jewelsMap := map[rune]int{}
	jewelsFounded := 0
	for _, jewel := range jewels {
		jewelsMap[jewel]++
	}

	for _, stone := range stones {
		if _, ok := jewelsMap[stone]; ok {
			jewelsFounded++
		}
	}

	return jewelsFounded
}

func ContainsDuplicate(nums []int) bool {
	dumps := map[int]bool{}

	for _, num := range nums {
		if _, ok := dumps[num]; ok {
			return true
		}

		dumps[num] = true
	}

	return false
}
