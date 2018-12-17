package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	tmp := nums[0]
	i, j := 1, 1
	for ; i < len(nums); i++ {
		for i < len(nums) && tmp == nums[i] {
			i++
		}
		if i >= len(nums) {
			break
		}
		nums[j] = nums[i]
		tmp = nums[i]
		j++
	}
	nums = nums[:j]
	fmt.Println(nums)
	return j
}

func main() {
	nums := []int{1, 1}
	fmt.Println(removeDuplicates(nums))
}
