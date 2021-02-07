package main

import (
"container/list"
"fmt"
)

// Binary Tree
type BinaryTree struct {
	Data  interface{}
	Val int
	Left  *BinaryTree
	Right *BinaryTree
}

// Constructor
func NewBinaryTree(data int) *BinaryTree {
	return &BinaryTree{Val: data}
}

// 先序遍历-非递归
func (bt *BinaryTree) PreOrderNoRecursion() []interface{} {
	t := bt
	stack := list.New()
	res := make([]interface{}, 0)
	for t != nil || stack.Len() != 0 {
		for t != nil {
			res = append(res, t.Data)//visit
			stack.PushBack(t)
			t = t.Left
		}
		if stack.Len() != 0 {
			v := stack.Back()
			t = v.Value.(*BinaryTree)
			t = t.Right
			stack.Remove(v)
		}
	}
	return res
}

// 中序遍历-非递归
func (bt *BinaryTree) InOrderNoRecursion() []interface{} {
	t := bt
	stack := list.New()
	res := make([]interface{}, 0)
	for t != nil || stack.Len() != 0 {
		for t != nil {
			stack.PushBack(t)
			t = t.Left

		}
		if stack.Len() != 0 {
			v := stack.Back()
			t = v.Value.(*BinaryTree)
			res = append(res, t.Data)//visit
			t = t.Right

			stack.Remove(v)
		}
	}
	return res
}
func inorderTraversalStack(root *BinaryTree) (res []int) {
	stack := []*BinaryTree{}
	for root != nil || len(stack) > 0 {
		//循环所有左节点
		for root != nil  {
			stack = append(stack, root)
			root = root.Left
		}
		//pop出最后一个节点
		root = stack[len(stack) -1]
		stack = stack[:len(stack)-1]
		//到res中
		res = append(res, root.Val)
		//将根节点置为右节点
		root = root.Right

		////循环找到所有左节点，存到stark中
		//for root != nil {
		//	//fmt.Println(root.Val)
		//	stack = append(stack, root)
		//	root = root.Left
		//}
		////根据最后一个做节点找最右节点
		//root = stack[len(stack)-1]
		//stack = stack[:len(stack)-1]
		//fmt.Println(root.Val)
		//
		//res = append(res, root.Val)
		//
		//root = root.Right


	}
	return
}

func inorderTraver(root *BinaryTree) (res []int){
	stack := []*BinaryTree{}

	for root!= nil ||  len(stack) >0{
		//循环到最左节点
		for root != nil  {
			stack = append(stack, root)
			root = root.Left
		}

		//弹出最后的节点到res
		root = stack[len(stack) -1]
		stack = stack[:len(stack) -1]
		res = append(res,root.Val)
		root = root.Right

	}
	return
}

func inorderTraversal(root *BinaryTree) (res []int) {
	for root != nil {
		if root.Left != nil {
			// predecessor 节点表示当前 root 节点向左走一步，
			// 然后一直向右走至无法走为止的节点
			predecessor := root.Left
			for predecessor.Right != nil && predecessor.Right != root {
				// 有右子树且没有设置过指向 root，则继续向右走
				predecessor = predecessor.Right
			}
			if predecessor.Right == nil {
				// 将 predecessor 的右指针指向 root，
				// 这样后面遍历完左子树 root.Left 后，就能通过这个指向回到 root
				predecessor.Right = root
				// 遍历左子树
				root = root.Left
			} else { // predecessor 的右指针已经指向了 root，
			// 则表示左子树 root.Left 已经访问完了
				res = append(res, root.Val)
				// 恢复原样
				predecessor.Right = nil
				// 遍历右子树
				root = root.Right
			}
		} else { // 没有左子树
			res = append(res, root.Val)
			// 若有右子树，则遍历右子树
			// 若没有右子树，则整颗左子树已遍历完，root 会通过之前设置的指向回到这颗子树的父节点
			root = root.Right
		}
	}
	return
}

// 后序遍历-非递归
func (bt *BinaryTree) PostOrderNoRecursion() []interface{} {
	t := bt
	stack := list.New()
	res := make([]interface{}, 0)
	var preVisited *BinaryTree

	for t != nil || stack.Len() != 0 {
		for t != nil {
			stack.PushBack(t)
			t = t.Left
		}

		v   := stack.Back()
		top := v.Value.(*BinaryTree)

		if (top.Left == nil && top.Right == nil) || (top.Right == nil && preVisited == top.Left) || preVisited == top.Right{
			res = append(res, top.Data)//visit
			preVisited = top
			stack.Remove(v)
		}else {
			t = top.Right
		}
	}
	return res
}
func preorderTraversal(root *BinaryTree ) (vals []int) {
	stack := []*BinaryTree{}
	node := root
	for node != nil || len(stack) > 0{
		for node != nil   {
			vals = append(vals,node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return
}
//func preorderTraversal(root *BinaryTree) (vals []int) {
//	stack := []*BinaryTree{}
//	node := root
	//for node != nil || len(stack) > 0 {
	//	for node != nil {
	//		vals = append(vals, node.Val)
	//		stack = append(stack, node)
	//		node = node.Left
	//	}
	//	node = stack[len(stack)-1].Right
	//	stack = stack[:len(stack)-1]
	//}


//	return
//}


func main() {
	t := NewBinaryTree(1)
	t.Left  = NewBinaryTree(3)
	t.Right = NewBinaryTree(6)
	t.Left.Left = NewBinaryTree(4)
	t.Left.Right = NewBinaryTree(5)
	t.Left.Left.Left = NewBinaryTree(7)

	//fmt.Println(t.PreOrderNoRecursion())
	fmt.Println(preorderTraversal(t))
	//fmt.Println(t.PostOrderNoRecursion())
}