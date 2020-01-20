package main

import "fmt"

/*
	题目：请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。
	例如，字符串“+100”，”5e2“，”-123“， “3.1416”及“-1E-16”都表示数值，
	但“12e”， “1a3.14”，“1.2.3”，“+-5”及“12e+5.4”都不是。
*/

func isNumeric(str *string) bool {
	if str == nil || len(*str) == 0 {
		return false
	}
	var numeric bool
	numeric = scanInteger(str)

	if len(*str) > 0 && (*str)[0] == '.' {
		*str = (*str)[1:]
		numeric = scanUnsignedInteger(str) || numeric
	}
	if len(*str) > 0 && ((*str)[0] == 'e' || (*str)[0] == 'E') {
		*str = (*str)[1:]
		numeric = numeric && len(*str) > 1 && scanInteger(str)
	}
	return numeric && len(*str) == 0
}

func scanUnsignedInteger(str *string) bool {
	count := 0
	for _, c := range *str {
		if c >= '0' && c <= '9' {
			count++
		} else {
			break
		}
	}
	*str = (*str)[count:]
	return count > 0
}

func scanInteger(str *string) bool {
	if (*str)[0] == '+' || (*str)[0] == '-' {
		*str = (*str)[1:]
	}
	return scanUnsignedInteger(str)
}

func main() {
	// fmt.Println(isNumeric("123"))
	// fmt.Println(isNumeric("-123"))
	// fmt.Println(isNumeric("-.14"))
	// fmt.Println(isNumeric("5e-2"))
	str := "1e.1"
	fmt.Println(isNumeric(&str))
	str = "5e-2"
	fmt.Println(isNumeric(&str))
	str = "121"
	fmt.Println(isNumeric(&str))
	str = "-1.34"
	fmt.Println(isNumeric(&str))
	str = "-1a34"
	fmt.Println(isNumeric(&str))
}
