package main

import (
	"fmt"
)

func rotate(nums []int, k int) {
	k %= len(nums)
	revert(nums)
	revert(nums[k:])
	revert(nums[:k])
}

func revert(nums []int) []int {
	length := len(nums)
	if length <= 1 {
		return nums
	}
	var tmp int
	for i := 0; i < length/2; i++ {
		tmp = nums[i]
		nums[i] = nums[length-1-i]
		nums[length-1-i] = tmp
	}
	return nums
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	fmt.Println(nums)

	rotate(nums, k)
	fmt.Println(nums)
}
