/**
 * @author zhangyuehao
 * @date 2019-08-12 16:58
 */

package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	t := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	fmt.Printf("res is %v", leverOrderTraversal(t))
}

func leverOrderTraversal(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var (
		res    [][]int
		temp   []int
		length int
		p      *TreeNode
	)

	stack := list.New()
	stack.PushFront(root)
	for stack.Len() != 0 {
		p = nil
		temp = make([]int, 0)
		length = stack.Len()
		for i := 0; i < length; i++ {
			e := stack.Back()
			stack.Remove(e)
			p = e.Value.(*TreeNode)
			if p.Left != nil {
				stack.PushFront(p.Left)
			}
			if p.Right != nil {
				stack.PushFront(p.Right)
			}

			temp = append(temp, p.Val)
		}
		res = append(res, temp)
	}
	return res
}
