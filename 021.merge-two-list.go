package leetcode

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	rnext := &ListNode{}
	result := &ListNode{
		Next: rnext}
	for l1 != nil || l2 != nil {
		if l1 == nil {
			rnext.Val = l2.Val
			rnext.Next = l2.Next
			return result.Next
		} else if l2 == nil {
			rnext.Val = l1.Val
			rnext.Next = l1.Next
			return result.Next
		}
		if l1.Val < l2.Val {
			rnext.Val = l1.Val
			l1 = l1.Next
			rnext.Next = &ListNode{}
			rnext = rnext.Next
		} else {
			rnext.Val = l2.Val
			l2 = l2.Next
			rnext.Next = &ListNode{}
			rnext = rnext.Next
		}
	}
	return result.Next
}
