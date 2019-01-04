package leetcode

var ans []string

func generateParenthesis(n int) []string {
	ans = nil
	backtrack("", 0, 0, n)
	return ans
}

func backtrack(cur string, open, close, max int) {
	if len(cur) == max*2 {
		ans = append(ans, cur)
		return
	}

	if open < max {
		backtrack(cur+"(", open+1, close, max)
	}
	if close < open {
		backtrack(cur+")", open, close+1, max)
	}
}
