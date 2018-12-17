package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	result := ""
	if len(strs) <= 0 {
		return result
	}
	str := strs[0]
	for index, value := range str {
		for i := 1; i < len(strs); i++ {
			if index >= len(strs[i]) || value != int32(strs[i][index]) {
				return result
			}
		}
		result += string(value)
	}
	return result
}

func main() {
	strs := []string{"flower", "flow", "flight"} // {"dog","racecar","car"}
	fmt.Println(strs)
	fmt.Println(longestCommonPrefix(strs))
}
