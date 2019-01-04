package leetcode

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	var ans int
	for i, j := 0, 0; j < len(s); j++ {
		if mi, ok := m[s[j]]; ok && mi > i {
			i = mi
		}
		ans = max(ans, j-i+1)
		m[s[j]] = j + 1
	}
	return ans
}
