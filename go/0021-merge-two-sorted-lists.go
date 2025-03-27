package main

//https://leetcode.com/problems/merge-two-sorted-lists/description/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	mergedHead := &ListNode{}

	current := mergedHead

	for list1 != nil && list2 != nil {

		var minVal int
		if list1.Val <= list2.Val {
			minVal = list1.Val
			list1 = list1.Next
		} else {
			minVal = list2.Val
			list2 = list2.Next
		}

		current.Next = &ListNode{Val: minVal}
		current = current.Next
	}

	for list1 != nil {
		current.Next = &ListNode{Val: list1.Val}
		list1 = list1.Next
		current = current.Next
	}

	for list2 != nil {
		current.Next = &ListNode{Val: list2.Val}
		list2 = list2.Next
		current = current.Next
	}

	return mergedHead.Next
}

/*
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	headMergedList := &ListNode{}

	currl1, currl2, currl3 := list1, list2, headMergedList

	for currl1 != nil || currl2 != nil {

		var node *ListNode

		if currl1 != nil && currl2 != nil {
			// minVal := min(currl1.Val, currl2.Val)

			var minVal int
			if currl1.Val <= currl2.Val {
				minVal = currl1.Val
				currl1 = currl1.Next
			} else {
				minVal = currl2.Val
				currl2 = currl2.Next
			}

			node = &ListNode{Val: minVal}

		} else {
			if currl1 != nil {
				node = &ListNode{Val: currl1.Val}
				currl1 = currl1.Next
			} else {
				node = &ListNode{Val: currl2.Val}
				currl2 = currl2.Next
			}
		}

		currl3.Next = node
		currl3 = currl3.Next
	}

	return headMergedList.Next

}
*/

// func main() {

// }
