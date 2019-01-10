package leetcode

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	f, s := head, head.Next

	for s != nil {
		f.Val, s.Val = s.Val, f.Val
		f = s.Next
		if f == nil {
			break
		}
		s = f.Next
	}
	return head
}
