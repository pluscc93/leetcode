package leetcode

import (
	"reflect"
	"testing"
)

func Test_mergeTwoList(t *testing.T) {
	l1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	l2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	r := mergeTwoLists(l1, l2)
	t1 := &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, nil}}}}}}
	if !reflect.DeepEqual(r, t1) {
		t.Fatalf("\n%+v \n%+v", r, t1)
	}

	r2 := mergeTwoLists(nil, nil)
	if r2 != nil {
		t.Fatalf("%+v", r2)
	}

}
