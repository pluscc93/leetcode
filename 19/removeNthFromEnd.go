package main

import "fmt"

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	pre := &ListNode{Next: head}
	first, second := pre, pre
	for i := 0; i <= n; i++ {
		first = first.Next
	}
	for first != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	return pre.Next
}

func main() {
	var nums []int
	for _, i := range nums {
		fmt.Println(i)
	}
	tmp := nums[0]
	fmt.Println(len(nums))
	fmt.Println(tmp)

}
