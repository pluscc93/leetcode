package leetcode

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_findMedianSortedArrays(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		want  float64
	}{
		{
			nums1: []int{1, 3},
			nums2: []int{2},
			want:  2.0,
		},
		{
			nums1: []int{1, 2},
			nums2: []int{3, 4},
			want:  2.5,
		},
		{
			nums1: []int{2, 3},
			nums2: []int{1},
			want:  2.0,
		},
	}

	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			got := findMedianSortedArrays(tc.nums1, tc.nums2)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("diff: (-want +got)\n%s", diff)
			}
		})
	}
}
