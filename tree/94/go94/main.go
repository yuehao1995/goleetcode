/**
 * @author zhangyuehao
 * @date 2019-08-12 14:21
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
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	fmt.Printf("res is %v", inOrderTraversal(t))
}

func inOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	p := root
	stack := list.New()
	for {
		e := stack.Back()
		if e == nil && p == nil {
			break
		}

		if p != nil {
			stack.PushBack(p)
			p = p.Left
			continue
		}
		stack.Remove(e)
		p, _ = e.Value.(*TreeNode)
		res = append(res, p.Val)
		p = p.Right

	}
	return res
}
