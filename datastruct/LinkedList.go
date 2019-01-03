package datastruct

import "reflect"

type ElementType interface{}

type ListNode struct {
	Val  ElementType
	Next *ListNode
}

type LinkedList struct {
	Head *ListNode
}

func NewListNode(x ElementType, next *ListNode) *ListNode {
	return &ListNode{x, next}
}

func NewLinkedList() *LinkedList {
	head := &ListNode{0, nil}
	return &LinkedList{head}
}

func (list *LinkedList) IsEmpty() bool {
	return list.Head.next == nil
}

func (list *LinkedList) Length() int {
	return int(reflect.ValueOf(list.Head.Val).Int())
}
