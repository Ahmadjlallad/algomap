package main

import (
	"github.com/Ahmadjlallad/algomap/internal/hashmapsAndSets"
	"testing"
)

func TestNumJewelsInStones(t *testing.T) {
	result := hashmapsAndSets.NumJewelsInStones("aA", "aAAbbbb")

	if result != 3 {
		t.Errorf("answer %d, expected %d", result, 3)
	}
}

func TestContainsDuplicate(t *testing.T) {
	result := hashmapsAndSets.ContainsDuplicate([]int{1, 2, 3, 1})
	if result != true {
		t.Errorf("answer %t, expected %t", result, true)
		return
	}

	result = hashmapsAndSets.ContainsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2})
	if result != false {
		t.Errorf("answer %t, expected %t", result, false)
	}
}
