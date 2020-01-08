package main

import (
	"fmt"
)

/*
	题目1：
	用两个栈实现一个队列。实现队列的appendTail和deleteHead，
	分别完成在队列尾部插入节点和在队列头部删除节点的功能

	题目2:
	用两个队列实现一个栈。
*/

type myQueue struct {
	s1 stack
	s2 stack
}

func (q *myQueue) appendTail(val int) {
	q.s1.push(val)
}

func (q *myQueue) deleteHead() int {
	if q.s2.len() <= 0 {
		for q.s1.len() > 0 {
			if val, ok := q.s1.pop(); ok {
				q.s2.push(val)
			}
		}
	}
	if q.s2.len() == 0 {
		fmt.Println("error")
	}
	val, _ := q.s2.pop()
	return val
}

type queue []int

func (l *queue) append(val int) {
	*l = append(*l, val)
}
func (l *queue) head() int {
	if l.len() > 0 {
		val := (*l)[0]
		*l = (*l)[1:]
		return val
	}
	return -1
}
func (l *queue) len() int {
	return len(*l)
}

type myStack struct {
	l1 queue
	l2 queue
}

func (s *myStack) push(val int) {
	if s.l1.len() > 0 {
		s.l1.append(val)
	} else if s.l2.len() > 0 {
		s.l2.append(val)
	} else {
		s.l1.append(val)
	}
}

func (s *myStack) pop() int {
	if s.l1.len() > 0 {
		for s.l1.len() > 1 {
			s.l2.append(s.l1.head())
		}
		return s.l1.head()
	} else if s.l2.len() > 0 {
		for s.l2.len() > 1 {
			s.l1.append(s.l2.head())
		}
		return s.l2.head()
	} else {
		fmt.Println("error")
	}
	return -1
}

func main() {
	q := myQueue{
		s1: stack{},
		s2: stack{},
	}
	q.appendTail(1)
	q.appendTail(2)
	fmt.Println(q.deleteHead())
	q.appendTail(3)
	fmt.Println(q.deleteHead())

	s := myStack{
		l1: queue{},
		l2: queue{},
	}
	s.push(1)
	s.push(2)
	s.push(3)
	fmt.Println(s.pop())
	fmt.Println(s.pop())
	fmt.Println(s.pop())
}
