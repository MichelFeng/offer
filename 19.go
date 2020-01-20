package main

/*
	题目：请实现一个函数用来匹配包含‘.’ 和 ‘*’ 的正则表达式。
	模式中的字符‘.’表示任意一个字符，而‘*’表示提签名的字符可以出现任意次（含0次）。
	在本体中，匹配是指字符串的所有字符匹配整个模式。例如，字符串“aaa”与模式“ab*ac*a”匹配，
	但与“aa.a”和“ab*a”均不匹配
*/

func match(str, pat string) bool {
	// 两个字符串都是空字符串
	if len(str) == 0 && len(pat) == 0 {
		return true
	}

	// str不为空，pat为空
	if len(str) != 0 && len(pat) == 0 {
		return false
	}

	// （str 为空，pat不为空） || （str不为空，pat不为空）

	// pat的第二个字符是‘*’的场景
	if len(pat) > 1 && pat[1] == '*' {
		// （str不为空，pat不为空）
		// pat的第一个字符与str的第一个字符相同或pat的第一个字符是‘.’
		if len(str) > 0 && (pat[0] == str[0] || pat[0] == '.') {
			// 判断新的字符串是否匹配
			return match(str[1:], pat[2:]) || match(str[1:], pat) || match(str, pat[2:])
		}
		//（str 为空，pat不为空）
		return match(str, pat[2:])
	}
	// 第二个字符不是‘*’的场景
	if len(str) > 0 && (str[0] == pat[0] || pat[0] == '.') {
		return match(str[1:], pat[1:])
	}
	return false
}

// func main() {
// 	fmt.Println(match("aaa", "ab*aa"))
// 	fmt.Println(match("aaa", "a.*b"))
// 	fmt.Println(match("aba", "ab*ac*a"))
// 	fmt.Println(match("aaa", "ab*ac*a"))
// }
