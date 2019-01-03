package leetcode

func twoSum(nums []int, target int) []int {
	arr := []int{-1, -1}
	for i := 0; i < len(nums); i++ {
		j := i + 1
		for j < len(nums) {
			sum := nums[i] + nums[j]
			if sum == target {
				arr[0] = i
				arr[1] = j
				return arr
			}
			j++
		}
	}
	return arr
}
