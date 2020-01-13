package main

/*
	题目：请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有字符的路径。
	路径可以从矩阵中的任意一格开始，每一步可以在矩阵中向左、右、上、下移动一格。
	如果一条路径经过了矩阵的某一格，那么该路径不能再次进入该格子。
*/

func hasPath(matrix []rune, rows, cols int, path []rune) bool {
	if len(matrix) == 0 || len(path) == 0 {
		return false
	}
	visited := make([]bool, rows*cols)
	var pathLength int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if hasPathCore(matrix, rows, cols, i, j, path, &pathLength, visited) {
				return true
			}
		}
	}
	return false
}

func hasPathCore(matrix []rune, rows, cols, row, col int, path []rune, pathLength *int, visited []bool) bool {
	if len(path) == *pathLength {
		return true
	}
	var hasPath bool
	if row >= 0 && row < rows && col >= 0 && col < cols && matrix[row*cols+col] == path[*pathLength] && !visited[row*cols+col] {
		*pathLength++
		visited[row*cols+col] = true
		// 向左
		l := hasPathCore(matrix, rows, cols, row, col-1, path, pathLength, visited)
		// 向上
		u := hasPathCore(matrix, rows, cols, row-1, col, path, pathLength, visited)
		// 向右
		r := hasPathCore(matrix, rows, cols, row, col+1, path, pathLength, visited)
		// 向下
		d := hasPathCore(matrix, rows, cols, row+1, col, path, pathLength, visited)
		hasPath = l || u || r || d
		if !hasPath {
			*pathLength--
			visited[row*cols+col] = false
		}
	}
	return hasPath
}

// func main() {
// 	fmt.Println(hasPath([]rune("abtgcfcsjdeh"), 3, 4, []rune("bfce")))
// 	fmt.Println(hasPath([]rune("abtgcfcsjdeh"), 3, 4, []rune("bdcg")))
// }
