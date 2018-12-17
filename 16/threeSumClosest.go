package main

import (
	"fmt"
)

func partition(nums []int, lo, hi int) int {
	val := nums[lo]
	for lo < hi {
		for lo < hi && nums[hi] >= val {
			hi--
		}
		nums[lo] = nums[hi]
		for lo < hi && nums[lo] <= val {
			lo++
		}
		nums[hi] = nums[lo]
	}
	nums[lo] = val
	return lo
}

// QuickSort 快速排序
func QuickSort(nums []int, lo, hi int) {
	if lo < hi {
		part := partition(nums, lo, hi)
		QuickSort(nums, lo, part-1)
		QuickSort(nums, part+1, hi)
	}
}

func threeSumClosest(nums []int, target int) int {
	QuickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
	var result int = 10000
	for i := 0; i < len(nums)-2; i++ {
		// 如果和前一个元素一样则跳过这个
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, len(nums)-1
		if nums[i]+nums[j]-target > abs(result-target) {
			fmt.Println(i, j, k, result)
			break
		}
		var sum int
		for j < k {
			sum = nums[i] + nums[j] + nums[k]
			if sum == target {
				return sum
			} else if sum > target {
				k--
			} else {
				j++
			}
			if abs(sum-target) < abs(result-target) {
				result = sum
			}
		}
	}
	return result
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
func main() {
	nums := []int{0, 1, 2} //{-1, 2, 1, -4}
	target := 0
	fmt.Println(threeSumClosest(nums, target))
}
