package main

import "fmt"

/*
面试题6 从尾到头打印链表

题目：输入一个链表的头节点，从尾到头反过来打印除每个节点的值。
*/

type node struct {
	val  int
	next *node
}

type stack []int

func (s *stack) push(val int) {
	*s = append(*s, val)
}

func (s *stack) pop() (int, bool) {
	if len(*s) == 0 {
		return 0, false
	}
	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return val, true
}

func (s *stack) len() int {
	return len(*s)
}

func traverseByStack(h *node) {
	var s = new(stack)
	if h == nil {
		return
	}
	n := h
	for n != nil {
		s.push(n.val)
		n = n.next
	}

	for i := s.len() - 1; i >= 0; i-- {
		if v, ok := s.pop(); ok {
			fmt.Println(v)
		}
	}
}

func traverseByRecursive(h *node) {
	if h != nil {
		if h.next != nil {
			traverseByRecursive(h.next)
		}
		fmt.Println(h.val)
	}
}

// func main() {
// 	h := node{}
// 	n1 := node{val: 1}
// 	n2 := node{val: 2}
// 	n3 := node{val: 3}
// 	n4 := node{val: 4}
// 	h.next = &n1
// 	n1.next = &n2
// 	n2.next = &n3
// 	n3.next = &n4

// 	traverseByStack(&h)

// 	traverseByRecursive(&h)
// }
