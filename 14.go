package main

/*
题目：给你一根长度为n的绳子，请把绳子剪成m段（m，n都是整数，n>1 并且 m>1）
	每段绳子的长度记为k[0],k[1]...k[m]。请问k[0]*k[1]*...*k[m]可能的最大乘积是多少？
	例如，当绳子的长度是8时，我们把它剪成长度为2、3、3的三段，此时得到的最大乘积是18.
*/

func maxProductAfterCutting(length int) int {
	if length < 2 {
		return 0
	}
	if length == 2 {
		return 1
	}
	if length == 3 {
		return 2
	}
	dp := make([]int, length+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3
	var max int
	for i := 4; i <= length; i++ {
		max = 0
		for j := 1; j <= i/2; j++ {
			product := dp[j] * dp[i-j]
			if max < product {
				max = product
			}
			dp[i] = max
		}
	}
	max = dp[length]
	return max
}

// func main() {
// 	fmt.Println(maxProductAfterCutting(5))
// 	fmt.Println(maxProductAfterCutting(8))
// }
