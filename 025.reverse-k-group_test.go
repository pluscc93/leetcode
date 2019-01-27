package leetcode

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	tests := []struct {
		lists *ListNode
		k     int
		want  *ListNode
	}{
		{
			lists: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			k:     2,
			want:  &ListNode{2, &ListNode{1, &ListNode{4, &ListNode{3, &ListNode{5, nil}}}}},
		},
		{
			lists: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			k:     3,
			want:  &ListNode{3, &ListNode{2, &ListNode{1, &ListNode{4, &ListNode{5, nil}}}}},
		},
		{
			lists: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			k:     1,
			want:  &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
		},
		{
			lists: &ListNode{1, &ListNode{2, nil}},
			k:     2,
			want:  &ListNode{2, &ListNode{1, nil}},
		},
	}

	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			got := reverseKGroup(tc.lists, tc.k)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("diff: (-want +got)\n%s", diff)
			}
		})
	}
}

func Test_revert(t *testing.T) {
	tests := []struct {
		lists *ListNode
		want  *ListNode
	}{
		{
			lists: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			want:  &ListNode{5, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{1, nil}}}}},
		},
	}

	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			got := revert(tc.lists)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("diff: (-want +got)\n%s", diff)
			}
		})
	}
}
