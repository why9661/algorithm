package leetcode

/*
删除链表的倒数第N个节点，并且返回链表的第一个节点

解：
此题中的head表示链表的第一个节点，而非头结点
如何找到倒数第n个节点（的前一个结点）:快慢指针
头结点（哨兵节点）：可以方便删除第一个元素。
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	result := &ListNode{}
	result.Next = head
	slow := result

	for head != nil {
		if n > 0 {
			head = head.Next
		} else {
			head = head.Next
			slow = slow.Next
		}
		n--
	}

	slow.Next = slow.Next.Next

	return result.Next
}
