package main

// https://leetcode.com/problems/two-sum/description/

func twoSum(nums []int, target int) []int {

	indexMap := make(map[int]int, len(nums)-1)

	for i, el := range nums {
		indexMap[el] = i
	}

	for i, el := range nums {
		diff := target - el
		value, isExist := indexMap[diff]

		if isExist && value != i {
			return []int{i, value}
		}
	}

	return nil
}

// func main() {

// 	nums := []int{2, 7, 11, 15}

// 	fmt.Println(twoSum(nums, 9))
// }
