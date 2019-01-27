package leetcode

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}
	nLen, n := len(needle), len(haystack)-len(needle)+1
	nStart := needle[0]
	for i := 0; i < n; i++ {
		if haystack[i] == nStart && haystack[i:i+nLen] == needle {
			return i
		}
	}
	return -1
}
