package main

import (
	"fmt"
)

func romanToInt(s string) int {
	fmt.Println("input:", s)
	result := 0
	scale := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'M':
			scale = 1000
		case 'D':
			scale = 500
		case 'C':
			scale = 100
			if i+1 < len(s) && s[i+1] == 'D' {
				scale = scale * 4
				i++
			}
			if i+1 < len(s) && s[i+1] == 'M' {
				scale = scale * 9
				i++
			}
		case 'L':
			scale = 50
		case 'X':
			scale = 10
			if i+1 < len(s) && s[i+1] == 'L' {
				scale = scale * 4
				i++
			}
			if i+1 < len(s) && s[i+1] == 'C' {
				scale = scale * 9
				i++
			}
		case 'V':
			scale = 5
		case 'I':
			scale = 1
			if i+1 < len(s) && s[i+1] == 'V' {
				scale = scale * 4
				i++
			}
			if i+1 < len(s) && s[i+1] == 'X' {
				scale = scale * 9
				i++
			}
		}
		result += scale
	}
	return result
}

func main() {
	arr := []string{"III", "IV", "IX", "LVIII", "MCMXCIV"}
	for _, value := range arr {
		fmt.Println(romanToInt(string(value)))
	}
}
