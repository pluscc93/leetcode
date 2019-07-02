package leetcode

func findSubstring(s string, words []string) []int {
	if len(s) == 0 || len(words) == 0 {
		return nil
	}
	res := []int{}
	var wLen int
	for i := 0; i < len(words); i++ {
		wLen += len(words[i])
	}

	for i := 0; i < len(s)-wLen+1; i++ {
		tmp := i
		m := make(map[string]int)
		for _, v := range words {
			m[v]++
		}
		flag := true
		for flag {
			flag = false
			sum := 0
			// 遍历map
			for k, v := range m {
				if v > 0 {
					if s[tmp:tmp+len(k)] == k {
						tmp = tmp + len(k)
						m[k]--
						flag = true
					}
					sum += m[k]
				}
			}
			if sum == 0 {
				res = append(res, i)
				break
			}
		}
	}
	return res
}
