package main

import "fmt"

/*
	题目：输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历的结果。
	如果是则返回true，否则返回false。假设输入的数组的任意两个数字都互不相同。
*/

func verifySequenceOfBST(sequence []int, length int) bool {
	if length <= 0 {
		return false
	}

	root := sequence[length-1]
	i := 0
	for ; i < length-1; i++ {
		if sequence[i] > root {
			break
		}
	}
	j := i
	for ; j < length-1; j++ {
		if sequence[j] < root {
			return false
		}
	}

	var left = true
	if i > 0 {
		left = verifySequenceOfBST(sequence[:i], i)
	}
	var right = true
	if i < length-1 {
		right = verifySequenceOfBST(sequence[i:], length-i-1)
	}
	return left && right

}

func main() {
	s := []int{7, 4, 6, 5}
	fmt.Println(verifySequenceOfBST(s, 4))

	s = []int{5, 7, 6, 9, 11, 10, 8}
	fmt.Println(verifySequenceOfBST(s, 7))
}
