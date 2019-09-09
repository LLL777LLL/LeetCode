package main

import "fmt"

/*给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
示例:
给定 1->2->3->4, 你应该返回 2->1->4->3.

*/

type listNode struct {
	Val int
	Next *listNode
}

//头插法
func addNode(n int,list *listNode) *listNode {
	node := &listNode{Val:n}
	node.Next = list
	return node
}

//尾插法
func appendNode(n int,list *listNode) *listNode {
	node := &listNode{Val:n}
	if list == nil {
		list = node
	}else {
		li := list
		for li.Next != nil {
			li = li.Next
		}
		li.Next = node
	}
	return list
}

func printLi(list *listNode) {
	var arr []int
	for list != nil {
		arr = append(arr,list.Val)
		list = list.Next
	}
	fmt.Println(arr)
}

//递归
func swapPairs(head *listNode) *listNode {
	if(head == nil || head.Next == nil){
		return head
	}
	next := head.Next
	head.Next = swapPairs(next.Next)
	next.Next = head
	return next
}

//链表的操作必须要确保整条链表中节点都还存在，比如下面的 指针 p ，在 n1 和 n2 交换之后，指针 p.Next 指向了 n2 节点，确保了链表的完整。
func swapPairs2(head *listNode) *listNode {
	node := &listNode{Val:0}
	node.Next = head
	p := node
	for p.Next != nil && p.Next.Next != nil {
		n1 := p.Next
		n2 := n1.Next

		n1.Next = n1.Next.Next
		n2.Next = n1
		p.Next = n2

		p = n1
	}
	return node.Next
}

func main() {
	//头插法
	var l1 *listNode
	//l1 = addNode(6,l1)
	l1 = addNode(5,l1)
	l1 = addNode(4,l1)
	l1 = addNode(3,l1)
	l1 = addNode(2,l1)
	l1 = addNode(1,l1)
	printLi(l1)
	printLi(swapPairs2(l1))
}
