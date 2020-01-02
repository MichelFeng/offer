package main

import "fmt"

/*
面试题4: 二维数组中的查找

在一个二维数组中，每一行都按照从左到右递增都顺序排序，
每一列都按照从上到下递增都顺序排序。请完成一个函数，
输入这样都一个二维数组和一个整数，判断数组中是否含有该整数。
*/
func find(matrix [][]int, number int) (exist bool) {
	if matrix == nil || len(matrix) == 0 {
		return false
	}

	r := len(matrix)
	c := len(matrix[0])
	for i := 0; i < r; i++ {
		for j := c - 1; j >= 0; j-- {
			if matrix[i][j] == number {
				return true
			} else if matrix[i][j] > number {
				continue
			} else {
				break
			}
		}
	}
	return false
}

func main() {
	matrix := [][]int{
		[]int{1, 2, 8, 9},
		[]int{2, 4, 9, 12},
		[]int{4, 7, 10, 13},
		[]int{6, 8, 11, 15},
	}
	fmt.Println(find(matrix, 5))
	fmt.Println(find(matrix, 8))
	fmt.Println(find(matrix, 16))
}
