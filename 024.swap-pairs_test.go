package leetcode

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_swapPairs(t *testing.T) {
	tests := []struct {
		lists *ListNode
		want  *ListNode
	}{
		{
			lists: &ListNode{2, &ListNode{1, &ListNode{4, &ListNode{3, nil}}}},
			want:  &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}},
		},
		{
			lists: &ListNode{1, nil},
			want:  &ListNode{1, nil},
		},
		{
			lists: &ListNode{1, &ListNode{2, nil}},
			want:  &ListNode{2, &ListNode{1, nil}},
		},
	}

	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			got := swapPairs(tc.lists)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("diff: (-want +got)\n%s", diff)
			}
		})
	}
}
