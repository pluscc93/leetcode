package main

import "fmt"

func twoSum(nums []int,target int)[]int {
	arr :=[]int{-1,-1}
	for i:=0;i<len(nums);i++ {
		j:=i+1
		for j<len(nums) {
			sum := nums[i]+nums[j]
			if sum == target {
				arr[0]=i
				arr[1]=j
				return arr
			}
			j++
		}
	}
	return arr
}

func main(){
	nums := [] int {2,7,11,15}
	target := 9
	arr:=twoSum(nums,target)
	for index,value := range arr {
		fmt.Println(index,value,nums[value])
	}
}