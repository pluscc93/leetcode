package leetcode

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{
			s:    "abcabcbb",
			want: 3,
		},
		{
			s:    "bbbbb",
			want: 1,
		},
		{
			s:    "pwwkew",
			want: 3,
		},
		{
			s:    "p",
			want: 1,
		},
		{
			s:    "abba",
			want: 2,
		},
	}

	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			got := lengthOfLongestSubstring(tc.s)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("diff: (-want +got)\n%s", diff)
			}
		})
	}
}
