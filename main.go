package main

import (
	"fmt"
	"strconv"
)

func main() {
}

// IsPalindrome Given an integer x, return true if x is a
// palindrome
// , and false otherwise.
func IsPalindrome(x int) bool {
	s := strconv.Itoa(x)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

// TwoSum takes an array of integers (`nums`) and a target integer (`target`).
// It returns a slice of integer slices, each containing a pair of indices.
// Each pair of indices corresponds to two distinct elements in `nums` that add up to `target`.
func TwoSum(nums []int, target int) [][]int {
	var s [][]int
	for i := 0; i < len(nums); i++ {
		fmt.Println("i ", i)

		for j := i + 1; j < len(nums); j++ {
			sum := nums[i] + nums[j]
			if sum == target {
				s = append(s, []int{i, j})
			}
			fmt.Println(j)

		}
	}
	return s
}
