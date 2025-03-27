package main

// https://leetcode.com/problems/add-two-numbers/description/

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
	Size int
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}

	curl3 := dummy

	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {

		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10
		curl3.Next = &ListNode{Val: sum % 10}
		curl3 = curl3.Next

	}

	return dummy.Next
}

func buildList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}
	head.Size++
	current := head

	for _, num := range nums[1:] {
		current.Next = &ListNode{Val: num}
		current = current.Next
		head.Size++
	}

	return head
}

func printList(head *ListNode) {
	current := head

	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
}

// func main() {

// 	nums1 := []int{2, 4, 3}
// 	nums2 := []int{5, 6, 4}

// 	l1 := buildList(nums1)
// 	l2 := buildList(nums2)

// 	l3 := addTwoNumbers(l1, l2)

// 	printList(l3)

// }
