package main

import "fmt"

/*
	题目：实现函数 double Power(double base, int exponent)，
	求base的exponent次方。不得使用库函数，同时不需要考虑大数问题。
*/

func power(base float64, exponent int) float64 {
	if base <= 1e-5 && exponent < 0 {
		// 0的负指数等价于0的正指数分支1，0的正指数是0，0的倒数没有意义
		fmt.Printf("invalid %v %v", base, exponent)
		return 0
	}

	absExponent := exponent
	if exponent < 0 {
		absExponent = -exponent
	}

	result := powerWithUnsignedExponent(base, absExponent)
	if exponent < 0 {
		result = 1.0 / result
	}
	return result
}

func powerWithUnsignedExponent(base float64, exponent int) float64 {
	// result := 1.0
	// for i := 1; i <= exponent; i++ {
	// 	result *= base
	// }
	// return result

	if exponent == 0 {
		// n的0次方等于1
		return 1
	}
	if exponent == 1 {
		// n的1次方等于n
		return base
	}
	// n的m次方 = （n 的m/2次方）的平方 * （1 if m为偶数 else n)
	result := powerWithUnsignedExponent(base, exponent>>1)
	result *= result
	if exponent&0x1 == 1 {
		result *= base
	}
	return result
}

// func main() {
// 	fmt.Println(power(2, 10))
// 	fmt.Println(power(2, -4))
// }
