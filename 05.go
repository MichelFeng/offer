package main

import (
	"fmt"
	"strings"
)

/*
面试题5: 替换空格

题目： 请实现一个函数，把字符串中的每个空格替换成%20.
例如，输入“We are happy.”，则输出“We%20are%20happy.”
*/

func replaceByBuiltin(s string) {
	fmt.Println(strings.Replace(s, " ", "%20", -1))
}

/*
相关题目：
有两个排序的数组A1和A2，内存在A1的末尾有足够多的空余空间容纳A2。
请实现一个函数，把A2中的所有数组插入A1中，并且所有的数字是排序的。

解题思路：
从后向前比较，较大的移动到A1的尾部。
*/

// func main() {
// 	s := "We are happy"
// 	replaceByBuiltin(s)
// }
