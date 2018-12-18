package main

import "fmt"

func nextPermutation(nums []int) {
	i := len(nums) - 2
	for i >= 0 && nums[i+1] <= nums[i] {
		i--
	}
	if i >= 0 {
		j := len(nums) - 1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		swap(nums, i, j)
	}
	revert(nums, i+1)
	return
}

func revert(nums []int, start int) {
	i, j := start, len(nums)-1
	for i < j {
		swap(nums, i, j)
		i++
		j--
	}
}

func swap(nums []int, i, j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}

func main() {
	nums := []int{1, 1}
	nextPermutation(nums)
	fmt.Println(nums)
}
