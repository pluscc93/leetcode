package leetcode

import "fmt"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var flag int
	result := &ListNode{0, nil}
	q := result
	for l1 != nil || l2 != nil {
		x := num(l1)
		y := num(l2)
		fmt.Print(x, y)
		q.Next = &ListNode{(x + y + flag) % 10, nil}
		fmt.Println(q.Next)
		flag = (x + y + flag) / 10
		q = q.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if flag != 0 {
		q.Next = &ListNode{flag, nil}
	}
	return result.Next
}

func num(l *ListNode) int {
	if l != nil {
		return l.Val
	}
	return 0
}
