package main

// https://leetcode.com/problems/median-of-two-sorted-arrays/description/

import (
	"slices"
)


func findMedianSortedArraysWithoutSort(nums1 []int, nums2 []int) float64 {

	var nums3 []int
	finalSize := len(nums1) + len(nums2)

	cur1, cur2 := 0, 0

	for i := 0; i < finalSize; i++ {
		if cur1 < len(nums1) && cur2 < len(nums2) {
			if nums1[cur1] < nums2[cur2] {
				nums3 = append(nums3, nums1[cur1])
				cur1++
			} else {
				nums3 = append(nums3, nums2[cur2])
				cur2++
			}
		} else {
			if cur1 >= len(nums1) {
				nums3 = append(nums3, nums2[cur2:]...)
			} else if cur2 >= len(nums2) {
				nums3 = append(nums3, nums1[cur1:]...)
			}
		}
	}

	n := finalSize - 1

	if n%2 == 0 {
		return float64(nums3[n/2])
	} else {
		return (float64(nums3[n/2] + nums3[n/2+1])) / 2
	}

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	nums1 = append(nums1, nums2...)

	slices.Sort(nums1)

	n := len(nums1) - 1
	if n%2 == 0 {
		return float64(nums1[n/2])
	} else {
		return (float64(nums1[n/2] + nums1[n/2+1])) / 2
	}
}

// func main() {

// 	nums1 := []int{1, 3}
// 	nums2 := []int{2, 4}

// 	median := findMedianSortedArrays(nums1, nums2)

// 	fmt.Println(median)
// }
