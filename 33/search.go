package main

import "fmt"

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	var mid int
	for left <= right {
		mid = (left + right) / 2
		fmt.Println(left, right, mid)
		fmt.Println(nums[left], nums[right], nums[mid])
		if nums[mid] == target {
			return mid
		} else if nums[mid] < nums[right] {
			if nums[mid] < target && nums[right] >= target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if nums[left] <= target && nums[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return -1
}

func main() {
	nums := []int{5, 1, 3}
	fmt.Println(search(nums, 5))
}
