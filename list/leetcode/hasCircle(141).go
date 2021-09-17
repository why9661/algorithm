package leetcode

/*
判断一个链表是否有环
*/

func hasCircle(head *ListNode) bool {
	if head == nil {
		return false
	}

	fast := head.Next
	slow := head

	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}

		fast = fast.Next.Next
		slow = slow.Next
	}

	return false
}
