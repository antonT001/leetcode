package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}

	node := new(ListNode)
	res := node

	for {
		if list1 == nil {
			node.Val = list2.Val
			node.Next = list2.Next
			return res
		}
		if list2 == nil {
			node.Val = list1.Val
			node.Next = list1.Next
			return res
		}

		if list1.Val <= list2.Val {
			node.Val = list1.Val
			list1 = list1.Next
		} else {
			node.Val = list2.Val
			list2 = list2.Next
		}

		if list1 == nil && list2 == nil {
			return res
		}
		node.Next = &ListNode{}
		node = node.Next
	}
}
