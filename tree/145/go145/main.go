/**
 * @author zhangyuehao
 * @date 2019-08-12 16:22
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
	fmt.Printf("res is %v", postOrderTraversal(t))
}

func postOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	p := root
	stack := list.New()

	var pLastVisit *TreeNode
	for p != nil {
		stack.PushBack(p)
		p = p.Left
	}

	for stack.Len() != 0 {
		e := stack.Back()

		pcur := e.Value.(*TreeNode)
		if pcur.Right == nil || pcur.Right == pLastVisit {
			res = append(res, pcur.Val)
			pLastVisit = pcur
			stack.Remove(e)
			continue
		}
		pcur = pcur.Right

		for pcur != nil {
			stack.PushBack(pcur)
			pcur = pcur.Left
		}
	}

	return res
}
