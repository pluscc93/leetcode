package main

import (
	"fmt"
)

var letterMap = map[int]string{
	2: "abc",
	3: "def",
	4: "ghi",
	5: "jkl",
	6: "mno",
	7: "pqrs",
	8: "tuv",
	9: "wxyz",
}

var result []string

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}
	result = nil
	combinations("", digits)
	return result
}

func combinations(letter, digits string) {
	if digits == "" {
		result = append(result, letter)
		return
	}

	key := int(digits[0] - '0')
	if v, ok := letterMap[key]; ok {
		for _, val := range v {
			combinations(letter+string(val), digits[1:])
		}
	}
}
func main() {
	// fmt.Println('0' - 48)
	fmt.Println(letterCombinations("22"))
	fmt.Println(letterCombinations(""))

}
