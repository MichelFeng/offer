package main

import "fmt"

/*
	题目：输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，
	本题从1开始计数，即链表的尾节点是倒数第1个节点。例如，一个链表有6个节点，
	从头节点开始，他们的值依次是1，2，3，4，5，6。这个链表的倒数第3个节点是
	值为4的节点。
*/

type listNodeFor22 struct {
	val  int
	next *listNodeFor22
}

func findKthToTail(head *listNodeFor22, k int) *listNodeFor22 {
	if k <= 0 || head == nil {
		return nil
	}
	ahead := head
	var behind *listNodeFor22 = nil

	for ; k > 0; k-- {
		if ahead != nil {
			ahead = ahead.next
		} else {
			return nil
		}
	}
	behind = head
	for ahead != nil {
		ahead = ahead.next
		behind = behind.next
	}
	return behind
}

// func main() {
// 	h := &listNodeFor22{val: 1}
// 	var t *listNodeFor22 = h
// 	for i := 2; i < 7; i++ {
// 		n := &listNodeFor22{val: i}
// 		t.next = n
// 		t = n
// 	}
// 	p := h
// 	for p != nil {
// 		fmt.Printf("%v\t", p.val)
// 		p = p.next
// 	}

// 	fmt.Printf("\n%v", findKthToTail(h, 3))
// }
