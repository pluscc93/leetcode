package leetcode

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_findSubstring(t *testing.T) {
	tests := []struct {
		s     string
		words []string
		want  []int
	}{
		{
			s:     "barfoothefoobarman",
			words: []string{"foo", "bar"},
			want:  []int{0, 9},
		},
		{
			s:     "wordgoodgoodgoodbestword",
			words: []string{"word", "good", "best", "word"},
			want:  []int{},
		},
		{
			s:     "",
			words: []string{"foo", "bar"},
			want:  nil,
		},
		{
			s:     "a",
			words: []string{},
			want:  nil,
		},
	}

	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			got := findSubstring(tc.s, tc.words)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("diff: (-want +got)\n%s", diff)
			}
		})
	}
}
