package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resultHead := &ListNode{}
	r := resultHead

	for {
		var v1, v2 int
		if l1 != nil {
			v1 = l1.Val
		}
		if l2 != nil {
			v2 = l2.Val
		}

		summ := r.Val + v1 + v2
		if summ <= 9 {
			r.Val = summ
		} else {
			r.Val = summ % 10
			r.Next = &ListNode{
				Val: summ / 10,
			}
		}

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}

		if l1 == nil && l2 == nil {
			return resultHead
		}

		if r.Next == nil {
			r.Next = &ListNode{}
		}
		r = r.Next

	}
}
