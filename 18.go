package main

/*
	题目一：在O(1)时间内删除链表节点
	给定单向链表的头指针和一个节点指针，定义一个函数在O(1)时间内删除该节点。
*/

type listNode struct {
	value int
	next  *listNode
}

func deleteNode(head, node *listNode) {
	if head == nil || node == nil {
		return
	}
	if node.next != nil {
		// 非尾节点
		n := node.next
		node.value = n.value
		node.next = n.next
		n = nil
	} else if head == node {
		// 只有一个节点
		head = nil
		node = nil
	} else {
		// 尾节点
		n := head
		for n.next != node {
			n = n.next
		}
		n.next = nil
		node = nil
	}
}

/*
	题目二：删除链表中重复的节点
	在一个排序的链表中，如何删除重复的节点？
*/
func deleteDuplication(head *listNode) {
	if head == nil {
		return
	}

	var pre *listNode
	node := head
	for node != nil {
		next := node.next
		var needDelete bool
		if next != nil && next.value == node.value {
			needDelete = true
		}
		if !needDelete {
			pre = node
			node = node.next
		} else {
			value := node.value
			del := node
			for del != nil && del.value == value {
				next = del.next
				del = nil
				del = next
			}
			if pre == nil {
				head = next
			} else {
				pre.next = next
			}
			node = next
		}
	}

}
// func main() {

// }
