package leetcode

// 使用快慢指针
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	result := &ListNode{}
	result.Next = head
	slow := result
	i := 0

	for head != nil {
		if i >= n {
			slow = slow.Next
		}
		head = head.Next
		i++
	}

	slow.Next = slow.Next.Next

	return result.Next
}
