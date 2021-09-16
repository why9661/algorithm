package leetcode

/*
两数相加
给你两个非空的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。
请你将两个数相加，并以相同形式返回一个表示和的链表。
你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	result := head
	temp := 0

	for l1 != nil || l2 != nil || temp != 0 {
		if l1 != nil {
			temp += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			temp += l2.Val
			l2 = l2.Next
		}

		newNode := &ListNode{Val: temp % 10}
		head.Next = newNode
		temp /= 10
		head = head.Next
	}

	return result.Next
}
