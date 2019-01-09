package leetcode

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	var r int
	r = lengthOfLongestSubstring("abcabcbb")
	if r != 3 {
		t.Fatal(r)
	}
	r = lengthOfLongestSubstring("bbbbb")
	if r != 1 {
		t.Fatal(r)
	}
	r = lengthOfLongestSubstring("pwwkew")
	if r != 3 {
		t.Fatal(r)
	}
	r = lengthOfLongestSubstring("p")
	if r != 1 {
		t.Fatal(r)
	}
	r = lengthOfLongestSubstring("abba")
	if r != 2 {
		t.Fatal(r)
	}
}
