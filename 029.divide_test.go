package leetcode

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_divide(t *testing.T) {
	tests := []struct {
		dividend int
		divisor  int
		want     int
	}{
		{
			dividend: 10,
			divisor:  3,
			want:     3,
		},
		{
			dividend: 7,
			divisor:  -3,
			want:     -2,
		},
		{
			dividend: 1,
			divisor:  1,
			want:     1,
		},
		{
			dividend: -2147483648,
			divisor:  -1,
			want:     2147483647,
		},
	}

	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			got := divide(tc.dividend, tc.divisor)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("diff: (-want +got)\n%s", diff)
			}
		})
	}
}
