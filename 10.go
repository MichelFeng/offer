package main

/*
	题目一： 求斐波那契数列的第n项
	写一个函数，输入n，求斐波那契数列的第n项
*/

// 递归方式
func fibonacci1(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci1(n-1) + fibonacci1(n-2)
}

// 遍历方式
func fibonacci2(n int) int {
	var (
		a = 0
		b = 1
	)
	if n <= 0 {
		return a
	}
	if n == 1 {
		return b
	}

	var c int
	for i := 2; i <= n; i++ {
		c = a + b
		a = b
		b = c
	}
	return c
}

/*
	题目二：青蛙跳台阶问题
	一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。
	求该青蛙跳上一个n级的台阶总共有多少种跳法。
*/
func jumpFrog(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return jumpFrog(n-1) + jumpFrog(n-2)
}

// func main() {
// 	fmt.Println(fibonacci1(5))
// 	fmt.Println(fibonacci2(5))

// 	fmt.Println(jumpFrog())
// }
