package leetcode

import (
	"reflect"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	lists := []*ListNode{
		&ListNode{1, &ListNode{4, &ListNode{5, nil}}},
		&ListNode{1, &ListNode{3, &ListNode{4, nil}}},
		&ListNode{2, &ListNode{6, nil}},
	}
	r := mergeKLists(lists)
	t1 := &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}}}
	if !reflect.DeepEqual(r, t1) {
		t.Fatalf("\n%+v \n%+v", r, t1)
	}
}
