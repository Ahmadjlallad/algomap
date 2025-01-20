package main

import (
	"testing"

	"github.com/Ahmadjlallad/algomap/internal/arrayAndString"
)

// TestClosest Given an integer array nums of size n, return the number with the value closest to 0 in nums. If there are multiple answers, return the number with the largest value.
// for a valid return value.
func TestClosest(t *testing.T) {
	examples := [][]int{
		{1, -4, -2, 1, 4, 8},
		{1, 2, -1, 1},
		{-100000, -100000, -100000},
	}
	for _, example := range examples {
		result := arrayAndString.FindClosestNumber(example[1:])
		if result != example[0] {
			t.Errorf("FindClosestNumber(%d) = %d, want %d", example[1:], result, example[0])
		}
	}
}

// TestMergeAlternately You are given two strings word1 and word2. Merge the strings by adding letters in alternating order, starting with word1. If a string is longer than the other, append the additional letters onto the end of the merged string.
func TestMergeAlternately(t *testing.T) {
	if result := arrayAndString.MergeAlternately("abc", "pqr"); result != "apbqcr" {
		t.Errorf("MergeAlternately(%s, %s) = %s, want %s", "abc", "pqr", result, "apbqcr")
	}

	if result := arrayAndString.MergeAlternately("ab", "pqrs"); result != "apbqrs" {
		t.Errorf("MergeAlternately(%s, %s) = %s, want %s", "ab", "pqrs", result, "apbqrs")
	}

	if result := arrayAndString.MergeAlternately("abcd", "pq"); result != "apbqcd" {
		t.Errorf("MergeAlternately(%s, %s) = %s, want %s", "abcd", "pq", result, "apbqcd")
	}
}

// TestRomanToInt Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
//
// Symbol       Value
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000
// For example, 2 is written as II in Roman numeral, just two ones added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.
//
// Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:
//
// I can be placed before V (5) and X (10) to make 4 and 9.
// X can be placed before L (50) and C (100) to make 40 and 90.
// C can be placed before D (500) and M (1000) to make 400 and 900.
// Given a roman numeral, convert it to an integer.
func TestRomanToInt(t *testing.T) {
	if result := arrayAndString.RomanToInt("III"); result != 3 {
		t.Errorf("RomanToInt(%s) = %d, want %d", "III", result, 3)
	}

	if result := arrayAndString.RomanToInt("LVIII"); result != 58 {
		t.Errorf("RomanToInt(%s) = %d, want %d", "LVIII", result, 58)
	}

	if result := arrayAndString.RomanToInt("MCMXCIV"); result != 1994 {
		t.Errorf("RomanToInt(%s) = %d, want %d", "MCMXCIV", result, 1994)
	}
}

func TestIsSubsequence(t *testing.T) {
	if result := arrayAndString.IsSubsequence("abc", "ahbgdc"); !result {
		t.Errorf("IsSubsequence(%s, %s) = %t, want %t", "abc", "ahbgdc", result, true)
	}

	if result := arrayAndString.IsSubsequence("axc", "ahbgdc"); result {
		t.Errorf("IsSubsequence(%s, %s) = %t, want %t", "axc", "ahbgdc", result, false)
	}

	if result := arrayAndString.IsSubsequence("b", "c"); result {
		t.Errorf("IsSubsequence(%s, %s) = %t, want %t", "b", "c", result, false)
	}

	if result := arrayAndString.IsSubsequence("c", "c"); !result {
		t.Errorf("IsSubsequence(%s, %s) = %t, want %t", "c", "c", result, true)
	}
}

// TestMaxProfit You are given an array prices where prices[i] is the price of a given stock on the ith day.
// You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
func TestMaxProfit(t *testing.T) {
	if result := arrayAndString.MaxProfit([]int{7, 1, 5, 3, 6, 4}); result != 5 {
		t.Errorf("MaxProfit([]int{7, 1, 5, 3, 6, 4}) = %d, want %d", result, 5)
	}

	if result := arrayAndString.MaxProfit([]int{7, 6, 4, 3, 1}); result != 0 {
		t.Errorf("MaxProfit([]int{7,6,4,3,1}) = %d, want %d", result, 0)
	}

	if result := arrayAndString.MaxProfit([]int{2, 4, 1}); result != 2 {
		t.Errorf("MaxProfit([]int{2,4,1}) = %d, want %d", result, 2)
	}

	if result := arrayAndString.MaxProfit([]int{2, 1, 2, 0, 1}); result != 1 {
		t.Errorf("MaxProfit([]int{2,1,2,0,1}) = %d, want %d", result, 1)
	}

	if result := arrayAndString.MaxProfit([]int{3, 2, 6, 5, 0, 3}); result != 4 {
		t.Errorf("MaxProfit([]int{3,2,6,5,0,3}) = %d, want %d", result, 4)
	}
}
