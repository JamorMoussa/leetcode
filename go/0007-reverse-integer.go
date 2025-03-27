package main

// https://leetcode.com/problems/reverse-integer/description/

import (
	"math"
)

func reverse(x int) int {

	reversedNumber := 0

	for x != 0 {
		reversedNumber = reversedNumber*10 + x%10
		x /= 10
	}

	if reversedNumber >= math.MaxInt32 || reversedNumber <= math.MinInt32 {
		return 0
	}

	return reversedNumber

}

// func main() {
// 	fmt.Println(reverse(-123))
// }
