package main

import "fmt"

//func maxArea(height []int) int {
//	if len(height) < 2 {
//		return 0
//	}
//	sum,tmpSum,tmpHi := 0,0,0
//	for index,hi := range height {
//		if hi < tmpHi {
//			continue
//		}
//		for i:=index+1;i<len(height);i++ {
//			h := height[i]
//			if hi < h {
//				h = hi
//			}
//			tmpSum = (i-index)*h
//			if tmpSum > sum {
//				sum = tmpSum
//				tmpHi = h
//			}
//		}	
//	}
//	return sum
//}

func maxArea(height []int) int {
	sum := 0
	lo,hi := 0,len(height)-1
	for lo < hi {
		h := height[lo]
		if h > height[hi] {
			h = height[hi] 
		}
		tmp := (hi-lo)*h
		if tmp > sum {
			sum = tmp
		}
		if height[lo] > height[hi] {
			hi--
		}else {
			lo++
		}
	}
	return sum
}

func main() {
	height := [] int {1,8,6,2,5,4,8,3,7}
	fmt.Println(maxArea(height))
}
