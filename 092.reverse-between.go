package leetcode

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil || m == n {
		return head
	}
	dummy := &ListNode{}
	dummy.Next = head
	tmp := dummy
	for i := 1; i < m; i++ {
		tmp = tmp.Next
	}
	pre := tmp.Next
	cur := tmp.Next.Next
	for i := m; i < n; i++ {
		pre.Next = cur.Next
		cur.Next = tmp.Next
		tmp.Next = cur
		cur = pre.Next
	}
	return dummy.Next
}
