package leetcode

/*
删除链表的倒数第N个节点，并且返回链表的第一个节点

解：
此题中的head表示链表的第一个节点，而非头结点
主要问题是如何找到倒数第n个节点（的前一个结点）？
1、循环一次得到链表长度，再循环找到目标结点
2、快慢指针
头结点（哨兵节点）：可以方便删除第一个元素。
*/
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
