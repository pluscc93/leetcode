package main

import (
	"container/list"
	"fmt"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	rStack := list.New()
	lStack := list.New()
	rStack.PushBack(root.Right)
	lStack.PushBack(root.Left)
	for rStack.Len() > 0 && lStack.Len() > 0 {
		r := rStack.Front()
		l := lStack.Front()
		rStack.Remove(r)
		lStack.Remove(l)
		rNode := r.Value.(*TreeNode)
		lNode := l.Value.(*TreeNode)
		if (rNode == nil && lNode != nil) || (rNode != nil && lNode == nil) {
			return false
		}
		if rNode != nil {
			if rNode.Val != lNode.Val {
				return false
			}
			rStack.PushBack(rNode.Right)
			rStack.PushBack(rNode.Left)
			lStack.PushBack(lNode.Left)
			lStack.PushBack(lNode.Right)
		}
	}
	if rStack.Len() != 0 || lStack.Len() != 0 {
		return false
	}
	return true
	// return isSymmetricV2(root.Left, root.Right)
}

func isSymmetricV2(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if (left == nil && right != nil) || (left != nil && right == nil) || (left.Val != right.Val) {
		return false
	}
	return isSymmetricV2(left.Left, right.Right) && isSymmetricV2(left.Right, right.Left)
}

// NewBinaryTree 新节点
func NewBinaryTree(val int) *TreeNode {
	return &TreeNode{Val: val}
}

// CreateBinaryTree 创建一颗二叉树
func CreateBinaryTree(arr []int) (t *TreeNode) {
	tmp := make(map[int]*TreeNode, len(arr))
	for index, val := range arr {
		fmt.Println(index, val)
		if val != 0 {
			tmp[index+1] = NewBinaryTree(val)
			if (index+1)%2 == 0 {
				if v, ok := tmp[(index+1)/2]; ok {
					v.Left = tmp[index+1]
				}
			} else {
				if v, ok := tmp[(index+1)/2]; ok {
					v.Right = tmp[index+1]
				}
			}
		}
	}
	if v, ok := tmp[1]; ok {
		t = v
	}
	fmt.Println(t)
	return t
}
func main() {
	arr := []int{1, 2, 2, 3, 4, 4, 3}
	input := CreateBinaryTree(arr)
	fmt.Println(isSymmetric(input))
}
