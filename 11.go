package main

/*
	题目：把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
	输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。
	例如，输入「3，4，5，1，2」为「1，2，3，4，5」的一个旋转，该数组的最小值为1
*/

func min(numbers []int) int {

	var (
		start = 0
		end   = len(numbers) - 1
		mid   = start
	)

	for numbers[start] >= numbers[end] {
		if end-start == 1 {
			mid = end
			break
		}

		mid = (start + end) / 2
		if numbers[start] == numbers[end] && numbers[start] == numbers[mid] {
			return minInOrder(numbers[start:end])
		}
		if numbers[mid] >= numbers[start] {
			start = mid
		} else if numbers[mid] <= numbers[end] {
			end = mid
		}
	}
	return numbers[mid]
}

func minInOrder(numbers []int) int {
	var result = numbers[0]
	for _, v := range numbers {
		if v < result {
			result = v
		}
	}
	return result
}

// func main() {
// 	fmt.Println(min([]int{1, 0, 1, 1, 1}))
// 	fmt.Println(min([]int{3, 4, 5, 1, 2}))
// }
