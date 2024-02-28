package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "()[{}]"
	fmt.Println(IsValid(s))
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

// RomanToInt Given a roman numeral, convert it to an integer.
func RomanToInt(s string) int {
	str := strings.ToUpper(s)

	romanNums := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0
	prevValue := 0
	for i := 0; i < len(str); i++ {
		value := romanNums[str[i]]
		if value > prevValue {
			total += value - 2*prevValue
		} else {
			total += value
		}
		prevValue = value
	}

	return total
}

// LongestCommonPrefix Given an array of strings return the longest common prefix string
// If there is no common prefix, return an empty string "".
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for _, str := range strs {
		for len(str) < len(prefix) || prefix != str[:len(prefix)] {
			prefix = prefix[:len(prefix)-1]

			if prefix == "" {
				return ""
			}
		}
	}

	return prefix
}

// IsValid Given a string s containing just the characters '(', ')', '{', '}', '[' and ']',
// determine if the input string is valid.
func IsValid(s string) bool {
	stack := []rune{}
	matchingBrackets := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}

	for _, char := range s {

		if opening, ok := matchingBrackets[char]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != opening {
				return false
			}

			stack = stack[:len(stack)-1]
		} else {

			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}
