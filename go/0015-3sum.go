package main

import "fmt"

// https://leetcode.com/problems/3sum/

func threeSum(nums []int) [][]int {

	triplets := [][]int{}

	for i := range len(nums) {

		for j := i + 1; j < len(nums); j++ {

			res := nums[i] + nums[j]
			for k := j + 1; k < len(nums); k++ {

				if res == -nums[k] {
					triplets = append(
						triplets, []int{nums[i], nums[j], nums[k]},
					)
				}
			}
		}
	}

	return triplets

}

func main() {

	list := []int{-1, 0, 1, 2, -1, -4}

	fmt.Println(threeSum(list))
}
