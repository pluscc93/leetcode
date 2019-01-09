package leetcode

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	for len(lists) > 1 {
		lists = append(lists, mergeTwoLists(lists[0], lists[1]))
		lists = lists[2:]
	}
	return lists[0]
}
