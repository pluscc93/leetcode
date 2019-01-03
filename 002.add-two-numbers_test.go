package leetcode

import (
	"fmt"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	var a = []int{2, 4, 3}
	var b = []int{5, 6, 4}
	var l1, l2 *ListNode
	l1 = &ListNode{a[0], nil}
	q := l1
	for i := 1; i < len(a); i++ {
		q.Next = &ListNode{a[i], nil}
		q = q.Next
	}
	k := l1
	for t != nil {
		fmt.Print(" -> ", k)
		k = k.Next
	}
	fmt.Println()
	l2 = &ListNode{b[0], nil}
	q = l2
	for i := 1; i < len(b); i++ {
		q.Next = &ListNode{b[i], nil}
		q = q.Next
	}
	k = l2
	for t != nil {
		fmt.Print(" -> ", k)
		k = k.Next
	}
	fmt.Println()
	result := addTwoNumbers(l1, l2)
	for result != nil {
		fmt.Print(" -> ", result)
		result = result.Next
	}
	fmt.Println(l2)
}
