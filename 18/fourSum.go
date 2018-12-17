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

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}
	result := [][]int{}
	var tmp []int
	QuickSort(nums, 0, len(nums)-1)
	fmt.Println(nums, tmp)
	for i := 0; i < len(nums)-3; i++ {
		// 如果和前一个元素一样则跳过这个
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			// 如果和前一个元素一样则跳过这个
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			k, l := j+1, len(nums)-1
			for k < l {
				sum := nums[i] + nums[j] + nums[k] + nums[l]
				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[k], nums[l]})
					for k < l && nums[k] == nums[k+1] {
						k++
					}
					for k < l && nums[l] == nums[l-1] {
						l--
					}
					k++
					l--
				} else if sum > target {
					l--
				} else {
					k++
				}
			}
		}
	}
	return result
}

func main() {
	nums := []int{-5, -4, -3, -2, -1, 0, 0, 1, 2, 3, 4, 5} // {0, 0, 0, 0} // {-3, -2, -1, 0, 0, 1, 2, 3}
	target := 0
	fmt.Println(fourSum(nums, target))
}
