package main

import (
	"fmt"
)

func primeFactory(num int) []int {
	var result []int
	flag := make([]bool, num)
	for i := 2; i <= num; i++ {
		// 不是质数跳过
		if flag[i] {
			continue
		}
		// 不是质数的标记为true
		for j := 2 * i; j < num; j += i {
			flag[j] = true
		}
		// 分解
		for num%i == 0 {
			result = append(result, i)
			num /= i
		}
	}
	return result
}

func main() {
	array := []int{435234}
	fmt.Println(primeFactory(array[0]))
}
