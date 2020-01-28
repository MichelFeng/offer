package main

import (
	"container/list"
	"fmt"
)

/*
	题目：从上往下打印除二叉树的每个节点，
	同一层的节点按照从左到右的顺序打印。
*/

type binaryTreeNode struct {
	val   int
	left  *binaryTreeNode
	right *binaryTreeNode
}

func printFromTopToBottom(root *binaryTreeNode) {
	if root == nil {
		return
	}
	l := list.List{}
	l.PushBack(root)
	for l.Len() > 0 {
		n := l.Front()
		l.Remove(n)
		node := n.Value.(*binaryTreeNode)
		fmt.Printf("%d\t", node.val)
		if node.left != nil {
			l.PushBack(node.left)
		}
		if node.right != nil {
			l.PushBack(node.right)
		}
	}
}

// func main() {
// 	n1 := &binaryTreeNode{val: 8}
// 	n2 := &binaryTreeNode{val: 6}
// 	n3 := &binaryTreeNode{val: 10}
// 	n1.left = n2
// 	n1.right = n3
// 	n4 := &binaryTreeNode{val: 5}
// 	n5 := &binaryTreeNode{val: 7}
// 	n2.left = n4
// 	n2.right = n5
// 	n6 := &binaryTreeNode{val: 9}
// 	n7 := &binaryTreeNode{val: 11}
// 	n3.left = n6
// 	n3.right = n7
// 	printFromTopToBottom(n1)

// }
