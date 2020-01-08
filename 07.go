package main

/*
	题目：输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。
	假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
	例如，输入前序遍历序列「1， 2， 4， 7， 3， 5， 6， 8」，
	和中序遍历序列「4， 7， 2， 1， 5， 3， 8， 6」，
	则重建二叉树并输出它的头节点。
*/
import (
	"fmt"
)

type treeNode struct {
	value int
	left  *treeNode
	right *treeNode
}

func constuct(preOrder, midOrder []int) *treeNode {

	if len(preOrder) == 0 || len(midOrder) == 0 {
		return nil
	}
	return constuctCore(preOrder, midOrder)
}

func constuctCore(preOrder, midOrder []int) *treeNode {
	root := &treeNode{}
	root.value = preOrder[0]

	var idx int
	for k, v := range midOrder {
		if v == preOrder[0] {
			idx = k
		}
	}

	if len(midOrder[:idx]) > 0 {
		root.left = constuctCore(preOrder[1:], midOrder[:idx])
	}

	if len(midOrder[idx+1:]) > 0 {
		root.right = constuctCore(preOrder[len(midOrder[:idx])+1:], midOrder[idx+1:])
	}

	return root
}

func traverse(root *treeNode) {
	preOrderTraverse(root)
	fmt.Println()
	midOrderTraverse(root)
	fmt.Println()
	postOrderTraverse(root)
	fmt.Println()
	orderTraverse(root)
}

func preOrderTraverse(root *treeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%v\t", root.value)
	if root.left != nil {
		preOrderTraverse(root.left)
	}
	if root.right != nil {
		preOrderTraverse(root.right)
	}
}

func midOrderTraverse(root *treeNode) {
	if root == nil {
		return
	}

	if root.left != nil {
		midOrderTraverse(root.left)
	}
	fmt.Printf("%v\t", root.value)
	if root.right != nil {
		midOrderTraverse(root.right)
	}
}

func postOrderTraverse(root *treeNode) {
	if root == nil {
		return
	}

	if root.left != nil {
		postOrderTraverse(root.left)
	}

	if root.right != nil {
		postOrderTraverse(root.right)
	}

	fmt.Printf("%v\t", root.value)
}

func orderTraverse(root *treeNode) {
	if root == nil {
		return
	}
	var s = make(chan *treeNode, 10)
	s <- root
	for {
		select {
		case n := <-s:
			fmt.Printf("%v\t", n.value)
			if n.left != nil {
				s <- n.left
			}
			if n.right != nil {
				s <- n.right
			}
		default:
			break
		}
	}

}

// func main() {
// 	preOrder := []int{1, 2, 4, 7, 3, 5, 6, 8}
// 	midOrder := []int{4, 7, 2, 1, 5, 3, 8, 6}

// 	head := constuct(preOrder, midOrder)
// 	traverse(head)
// }
