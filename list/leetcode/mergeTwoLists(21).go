package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	result := head
	for l1.Next != nil && l2.Next != nil {
		if l1.Val < l2.Val {
			head.Next = l1
			l1 = l1.Next
		} else {
			head.Next = l2
			l2 = l2.Next
		}
		head = head.Next
	}

	if l1.Next == nil {
		head.Next = l2
	}

	if l2.Next == nil {
		head.Next = l1
	}

	return result.Next
}
