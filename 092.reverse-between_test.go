package leetcode

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_reverseBetween(t *testing.T) {
	tests := []struct {
		lists *ListNode
		m     int
		n     int
		want  *ListNode
	}{
		{
			lists: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			m:     2,
			n:     4,
			want:  &ListNode{1, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{5, nil}}}}},
		},
		{
			lists: &ListNode{3, &ListNode{5, nil}},
			m:     1,
			n:     2,
			want:  &ListNode{5, &ListNode{3, nil}},
		},
		{
			lists: &ListNode{3, nil},
			m:     1,
			n:     1,
			want:  &ListNode{3, nil},
		},
	}

	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			got := reverseBetween(tc.lists, tc.m, tc.n)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("diff: (-want +got)\n%s", diff)
			}
		})
	}
}
