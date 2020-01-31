package main

import "fmt"

/*
	题目：请实现函数 ComplexListNode* Clone(ComplexListNode* pHead),
	复制一个复杂链表。在复杂链表中，每个节点除了有一个m_pNext 指针指向下一个节点外，
	还有一个m_pSibling 指向链表中的任意节点或者NULL。
*/

type complexListNode struct {
	val     int
	next    *complexListNode
	sibling *complexListNode
}

func cloneNodes(head *complexListNode) {
	node := head
	for node != nil {
		cloned := &complexListNode{}
		cloned.val = node.val
		cloned.next = node.next

		node.next = cloned
		node = cloned.next
	}
}

func connectSiblingNodes(head *complexListNode) {
	node := head
	for node != nil {
		cloned := node.next
		if node.sibling != nil {
			cloned.sibling = node.sibling.next
		}
		node = cloned.next
	}
}

func reconnectNodes(head *complexListNode) *complexListNode {
	node := head
	var (
		clonedHead *complexListNode
		clonedNode *complexListNode
	)

	if node != nil {
		clonedHead = node.next
		clonedNode = node.next

		node.next = clonedNode.next
		node = node.next
	}

	for node != nil {
		clonedNode.next = node.next
		clonedNode = clonedNode.next
		node.next = clonedNode.next
		node = node.next
	}
	return clonedHead
}

func clone(head *complexListNode) *complexListNode {
	cloneNodes(head)
	connectSiblingNodes(head)
	return reconnectNodes(head)
}

func main() {
	n1 := &complexListNode{val: 1}
	n2 := &complexListNode{val: 2}
	n3 := &complexListNode{val: 3}
	n4 := &complexListNode{val: 4}
	n5 := &complexListNode{val: 5}
	n1.next = n2
	n1.sibling = n3
	n2.next = n3
	n2.sibling = n5
	n3.next = n4
	n4.next = n5
	n4.sibling = n2

	cloned := clone(n1)
	for cloned != nil {
		fmt.Printf("%d %v\t", cloned.val, cloned.sibling)
		cloned = cloned.next
	}
}
