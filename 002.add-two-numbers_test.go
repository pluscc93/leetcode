package leetcode

import (
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	r := addTwoNumbers(l1, l2)
	if r.Val != 7 {
		t.Fatalf("%+v", r)
	}
	r = r.Next
	if r.Val != 0 {
		t.Fatalf("%+v", r)
	}
	r = r.Next
	if r.Val != 8 {
		t.Fatalf("%+v", r)
	}
}
