package main

/*
	题目：输入一个整数数组，实现一个函数来调整该数组中数字的顺序，
	使得所有奇数位于数组的前半部分，所有偶数位于数组的后半部分。
*/

func reorder(numbers []int, f func(x int) bool) {
	if len(numbers) == 0 {
		return
	}
	start := 0
	end := len(numbers) - 1
	for start < end {
		for start < end && !f(numbers[start]) {
			start++
		}
		for start < end && f(numbers[end]) {
			end--
		}
		if start < end {
			numbers[start], numbers[end] = numbers[end], numbers[start]
		}
	}
}

func isEven(number int) bool {
	return number&1 == 0
}

// func main() {
// 	numbers := []int{1, 2, 3, 4, 5}
// 	reorder(numbers, isEven)
// 	fmt.Println(numbers)
// }
