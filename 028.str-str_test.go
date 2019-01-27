package leetcode

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_strStr(t *testing.T) {
	tests := []struct {
		haystack string
		needle   string
		want     int
	}{
		{
			haystack: "hello",
			needle:   "ll",
			want:     2,
		},
		{
			haystack: "aaaaa",
			needle:   "bba",
			want:     -1,
		},
		{
			haystack: "hello",
			needle:   "",
			want:     0,
		},
		{
			haystack: "a",
			needle:   "a",
			want:     0,
		},
	}

	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			got := strStr(tc.haystack, tc.needle)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("diff: (-want +got)\n%s", diff)
			}
		})
	}
}
