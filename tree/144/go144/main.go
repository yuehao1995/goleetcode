/**
 * @author zhangyuehao
 * @date 2019-08-12 11:29
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
	fmt.Printf("res is %v", preOrderTraversal(t))
}

//func preOrderTraversal(root *TreeNode) []int {
//   if root==nil{
//   	return nil
//   }
//   res:=[]int{root.Val}
//   leftRes := preOrderTraversal(root.Left)
//   rightRes := preOrderTraversal(root.Right)
//   res=append(res,leftRes...)
//   res = append(res, rightRes...)
//
//   return res
//}

func preOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	stack := list.New()
	stack.PushBack(root)
	for {
		e := stack.Back()
		if e == nil {
			break
		}
		stack.Remove(e)
		t, _ := e.Value.(*TreeNode)
		res = append(res, t.Val)

		if t.Right != nil {
			stack.PushBack(t.Right)
		}
		if t.Left != nil {
			stack.PushBack(t.Left)
		}
	}
	return res
}
