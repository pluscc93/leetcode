package leetcode

import (
	"testing"
)

func Test_mergeTwoList(t *testing.T) {
	l1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	l2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	r := mergeTwoLists(l1, l2)
	if r.Val != 1 {
		t.Fatalf("1:%+v", r)
	}
	r = r.Next
	if r.Val != 1 {
		t.Fatalf("2:%+v", r)
	}
	r = r.Next
	if r.Val != 2 {
		t.Fatalf("3:%+v", r)
	}
	r = r.Next
	if r.Val != 3 {
		t.Fatalf("4:%+v", r)
	}
	r = r.Next
	if r.Val != 4 {
		t.Fatalf("5:%+v", r)
	}
	r = r.Next
	if r.Val != 4 {
		t.Fatalf("6:%+v", r)
	}

	r2 := mergeTwoLists(nil, nil)
	if r2 != nil {
		t.Fatalf("%+v", r2)
	}

}
