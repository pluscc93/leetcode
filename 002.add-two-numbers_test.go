package leetcode

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	tests := []struct {
		l1   *ListNode
		l2   *ListNode
		want *ListNode
	}{
		{
			l1:   &ListNode{2, &ListNode{4, &ListNode{3, nil}}},
			l2:   &ListNode{5, &ListNode{6, &ListNode{4, nil}}},
			want: &ListNode{7, &ListNode{0, &ListNode{8, nil}}},
		},
	}

	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			got := addTwoNumbers(tc.l1, tc.l2)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("diff: (-want +got)\n%s", diff)
			}
		})
	}
}
