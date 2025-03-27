package main

// https://leetcode.com/problems/palindrome-number/description/

import (
	"strconv"
)

func isPalindromeStrDef(x int) bool {

	if x < 0 {
		return false
	}

	num_s := strconv.Itoa(x)
	length := len(num_s)

	if length == 2 {
		return x%11 == 0
	}

	for i := range length / 2 {
		if num_s[i] != num_s[length-i-1] {
			return false
		}
	}

	return true
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	reversed := 0
	original := x
	for x > 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}

	return reversed == original
}

// func main() {

// 	num := 1101

// 	fmt.Println(isPalindrome(num))

// }
