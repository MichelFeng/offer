package main

import "fmt"

/*
	题目：输入数字n，按顺序打印出从1到最大的n位十进制数。
	比如输入3，则打印出1，2，3一直到最大的3位数999
*/

func print1ToMaxOfNDigits(n int) {
	if n <= 0 {
		return
	}
	number := make([]rune, n)
	for i := 0; i < 10; i++ {
		number[0] = rune(i + int('0'))
		print1ToMaxOfNDigitsRecursively(number, n, 0)
	}
}

func print1ToMaxOfNDigitsRecursively(number []rune, length, index int) {
	if index == length-1 {
		printNumber(number)
		return
	}

	for i := 0; i < 10; i++ {
		number[index+1] = rune(i + int('0'))
		print1ToMaxOfNDigitsRecursively(number, length, index+1)
	}
}

func printNumber(number []rune) {
	isBeginning0 := true
	length := len(number)
	for i := 0; i < length; i++ {
		if isBeginning0 && number[i] != '0' {
			isBeginning0 = false
		}
		if !isBeginning0 {
			fmt.Printf("%v", string(number[i]))
		}
	}
	fmt.Print("\t")
}

// func main() {
// 	print1ToMaxOfNDigits(3)
// 	fmt.Println()
// 	// print1ToMaxOfNDigits(10)
// }
