package leetcode

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	r := addTwoNumbers(l1, l2)
	t1 := &ListNode{7, &ListNode{0, &ListNode{8, nil}}}
	if !reflect.DeepEqual(r, t1) {
		t.Fatalf("\n%+v \n%+v", r, t1)
	}
}
