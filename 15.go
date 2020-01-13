package main

import "fmt"

/*
	题目：请实现一个函数，输入一个整数，输出该数二进制表示中1的个数。
	例如，把9表示成二进制是1001，有2位是1.因此，如果输入9，则该函数输出2.
*/

func numOf1(n int) int {
	var count int
	for n > 0 {
		n = n & (n - 1)
		count++
	}
	return count
}
func main() {
	fmt.Println(numOf1(9))
}
