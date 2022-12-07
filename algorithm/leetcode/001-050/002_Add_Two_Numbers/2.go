package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	temp := &ListNode{}
	current := temp
	carry := 0
	for l1 != nil || l2 != nil {
		val1, val2 := 0, 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}
		sum := val1 + val2 + carry
		current.Next = &ListNode{Val: sum % 10}
		carry = sum / 10
		current = current.Next
	}
	if carry != 0 {
		current.Next = &ListNode{Val: carry}
	}
	return temp.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
