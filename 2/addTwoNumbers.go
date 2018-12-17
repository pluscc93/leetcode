package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode{
	var flag int = 0
	result := &ListNode{0,nil}
	q:=result
	for l1!=nil || l2!=nil {
		x:=num(l1)
		y:=num(l2)
		fmt.Print(x,y)
		q.Next=&ListNode{(x+y+flag)%10,nil}
		fmt.Println(q.Next)
		flag=(x+y+flag)/10
		q=q.Next
		if l1!=nil {
			l1 = l1.Next
		}
		if l2!=nil {
			l2 = l2.Next
		}
	}
	if flag!=0 {
		q.Next=&ListNode{flag,nil}
	}
	return result.Next
}

func num(l *ListNode) int {
	if l!=nil {
		return l.Val
	}
	return 0
}
func main() {
	var a = []int {2,4,3}
	var b = []int {5,6,4}
	var l1,l2 *ListNode
	l1 = &ListNode{a[0],nil}
	q:=l1
	for i:=1;i<len(a);i++ {
		q.Next=&ListNode{a[i],nil}
		q=q.Next
	}
	t:=l1
	for t!= nil {
		fmt.Print(" -> ",t)
		t=t.Next
	}
	fmt.Println()
	l2 = &ListNode{b[0],nil}
	q=l2
	for i:=1;i<len(b);i++ {
		q.Next=&ListNode{b[i],nil}
		q=q.Next
	}
	t=l2
	for t!= nil {
		fmt.Print(" -> ",t)
		t=t.Next
	}
	fmt.Println()
	result := addTwoNumbers(l1,l2)
	for result != nil {
		fmt.Print(" -> ",result)
		result=result.Next
	}
	fmt.Println(l2)
}
