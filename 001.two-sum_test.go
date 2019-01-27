package leetcode

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_twoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
	}

	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			got := twoSum(tc.nums, tc.target)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("diff: (-want +got)\n%s", diff)
			}
		})
	}
}
