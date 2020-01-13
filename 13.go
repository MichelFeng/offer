package main

/*
	题目：地上有一个m行n列的方格。
	一个机器人从坐标(0,0)的格子开始移动，它每次可以向左、右、上、下移动一格，
	但不能进入行坐标和列坐标但数位之和大于k的格子。
	例如，当k为18时，机器人能够进入方格(35,37)，因为3+5+3+7=18。但它不能进入方格(35,38)。
	请问该机器人能够到达多少个格子？
*/

func movingCount(threshold, rows, cols int) int {
	if threshold < 0 || rows <= 0 || cols <= 0 {
		return 0
	}
	visited := make([]bool, cols*rows)
	return movingCountCore(threshold, rows, cols, 0, 0, visited)
}

func movingCountCore(threshold, rows, cols, row, col int, visited []bool) int {
	var count int
	if check(threshold, rows, cols, row, col, visited) {
		visited[row*cols+col] = true

		count = 1 + movingCountCore(threshold, rows, cols, row-1, col, visited) + movingCountCore(threshold, rows, cols, row, col-1, visited) + movingCountCore(threshold, rows, cols, row+1, col, visited) + movingCountCore(threshold, rows, cols, row, col+1, visited)
	}
	return count
}

func check(threshold, rows, cols, row, col int, visited []bool) bool {
	if row >= 0 && row < rows && col >= 0 && col < cols && getDigitSum(row)+getDigitSum(col) <= threshold && !visited[row*cols+col] {
		return true
	}
	return false
}

func getDigitSum(num int) int {
	var sum int
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}

// func main() {
// 	fmt.Println(movingCount(5, 4, 4))
// }
