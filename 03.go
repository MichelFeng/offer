package main

// 面试题3 数组中重复的数字
/*
题目1: 在一个长度为n的数组里的所有数字都在 0～n-1 的范围内。
数组中的某些数字是重复的，但不知道有几个数字重复了，
也不知道没个数字重复了几次，请找出数组中任意一个重复的数字。
*/
func duplicate(numbers []int) (num int, hasDuplicate bool) {
	if numbers == nil || len(numbers) == 0 {
		return
	}

	max := len(numbers) - 1

	for i, v := range numbers {
		if v < 0 || v >= max {
			return
		}

		for v != i {
			if v == numbers[v] {
				return v, true
			}
			tmp := numbers[v]
			numbers[v] = numbers[i]
			numbers[i] = tmp
		}
	}
	return
}

/*
题目2: 不修改数组找出重复的数字
在一个长度为n+1的数组里的所有数字都在1～n的范围内，所以数组中至少有一个数字是重复的。
请找出数组中任意一个重复的数字，但不能修改输入的数组。例如输入长度为8的数组{2,3,5,4,3,2,6,7}，
那么对应的输出是重复的数字2或者3.
*/
func getDupliation(numbers []int) (num int, hasDupliate bool) {
	if numbers == nil || len(numbers) == 0 {
		return
	}
	start := 1
	end := len(numbers) - 1
	for start <= end {
		middle := (end-start)>>1 + start
		count := countRange(numbers, start, middle)
		if start == end {
			if count > 1 {
				return start, true
			}
			break
		}
		if count > (middle - start + 1) {
			end = middle
		} else {
			start = middle + 1
		}
	}
	return
}

func countRange(numbers []int, start, end int) (count int) {
	for _, v := range numbers {
		if v >= start && v <= end {
			count++
		}
	}
	return
}

// func main() {
// 	// fmt.Println(duplicate([]int{2, 3, 1, 0, 2, 5, 3}))
// 	fmt.Println(getDupliation([]int{2, 3, 1, 4, 2, 5, 3}))
// 	fmt.Println(getDupliation([]int{2, 3, 5, 7, 8, 2, 5, 3}))
// }
