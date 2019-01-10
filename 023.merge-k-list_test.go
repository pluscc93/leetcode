package leetcode

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	tests := []struct {
		lists []*ListNode
		want  *ListNode
	}{
		{
			lists: []*ListNode{
				&ListNode{1, &ListNode{4, &ListNode{5, nil}}},
				&ListNode{1, &ListNode{3, &ListNode{4, nil}}},
				&ListNode{2, &ListNode{6, nil}},
			},
			want: &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}}},
		},
		{
			lists: []*ListNode{nil, nil},
			// want:  &ListNode{},
		},
	}

	for _, tc := range tests {
		got := mergeKLists(tc.lists)
		if diff := cmp.Diff(tc.want, got); diff != "" {
			t.Errorf("diff: (-want +got)\n%s", diff)
		}
	}
}
