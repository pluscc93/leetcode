package leetcode

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{}
	tmp := head
	var count int
	for tmp != nil {
		count++
		tmp = tmp.Next
	}

	dummy.Next = head
	pre := head
	cur := head.Next
	l := dummy
	for count >= k {
		for i := 1; i < k; i++ {
			pre.Next = cur.Next
			cur.Next = l.Next
			l.Next = cur
			cur = pre.Next
		}
		if cur == nil {
			break
		}
		l = pre
		pre = pre.Next
		cur = cur.Next
		count -= k
	}
	return dummy.Next
}

func revert(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{}
	dummy.Next = head
	pre := head
	cur := head.Next
	for cur != nil {
		pre.Next = cur.Next
		cur.Next = dummy.Next
		dummy.Next = cur
		cur = pre.Next
	}
	return dummy.Next
}
