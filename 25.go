package main

import (
	"fmt"
)

/*
	题目：输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入的整数的所有路径。
	从树的根节点开始往下一直到叶子结点所经过的节点形成一条路径。
*/

type binaryTreeNode25 struct {
	val   int
	left  *binaryTreeNode25
	right *binaryTreeNode25
}

func findPath(root *binaryTreeNode25, expectedSum int) {
	if root == nil {
		return
	}
	var path []*binaryTreeNode25
	currentSum := 0
	findPathRecursively(root, expectedSum, path, currentSum)
}

func findPathRecursively(root *binaryTreeNode25, expectedSum int, path []*binaryTreeNode25, currentSum int) {
	currentSum += root.val
	path = append(path, root)

	isLeaf := root.left == nil && root.right == nil
	if expectedSum == currentSum && isLeaf {
		fmt.Printf("a path is found: ")
		for _, n := range path {
			fmt.Printf("%d\t", n.val)
		}
		fmt.Println()
	}
	if root.left != nil {
		findPathRecursively(root.left, expectedSum, path, currentSum)
	}
	if root.right != nil {
		findPathRecursively(root.right, expectedSum, path, currentSum)
	}
	currentSum -= root.val
	l := len(path)
	if l >= 2 {
		path = path[:len(path)-2]
	}
}
func main() {
	n1 := &binaryTreeNode25{val: 10}
	n2 := &binaryTreeNode25{val: 5}
	n3 := &binaryTreeNode25{val: 12}
	n4 := &binaryTreeNode25{val: 4}
	n5 := &binaryTreeNode25{val: 7}
	n1.left = n2
	n1.right = n3
	n2.left = n4
	n2.right = n5

	findPath(n1, 22)
}
