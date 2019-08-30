package main

import (
	"fmt"
)

/*将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 
示例：输入：1->2->4, 1->3->4   输出：1->1->2->3->4->4  ;
*/

//共用
type ListNode struct {
	Val int
	Next *ListNode
}



//ListNode 头插法
func AddListNode(num int,no *ListNode) *ListNode {
	li := &ListNode{Val:num}
	li.Next = no
	no = li
	return no
}

//用递归的方式实现
func mergeTwoLists(l1,l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var node *ListNode
	if l1.Val <  l2.Val {
		node = l1
		node.Next = mergeTwoLists(l1.Next,l2)
	}else {
		node = l2
		node.Next = mergeTwoLists(l2.Next,l1)
	}
	return node
}





type List struct {
	headNode *ListNode
}

//头插法
func (list *List) Add(num int) {
	node := &ListNode{Val:num}
	node.Next = list.headNode
	list.headNode = node
}


//最常规的思想实现，一般最常规的思想实现代码都比较臃肿，考虑用用 递归
func mergeTwoLists2(l11,l22 List) *ListNode {
	l1 , l2 := l11.headNode,l22.headNode
	var node *ListNode
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			no := &ListNode{Val:l1.Val}
			if node == nil {
				node = no
			}else {
				li := node
				for li.Next != nil {
					li = li.Next
				}
				li.Next = no
			}
			l1 = l1.Next
		}else {
			no := &ListNode{Val:l2.Val}
			if node == nil {
				node = no
			}else {
				li := node
				for li.Next != nil {
					li = li.Next
				}
				li.Next = no
			}
			l2 = l2.Next
		}
	}
	if l1 == nil && l2 != nil {
		li := node
		if li == nil {
			node = l2
		}else {
			for li.Next != nil {
				li = li.Next
			}
			li.Next = l2
		}
	}else if l1 != nil && l2 ==nil {
		li := node
		if li == nil {
			node = l1
		}else {
			for li.Next != nil {
				li = li.Next
			}
			li.Next = l1
		}
	}
	return node
}

func P(node *ListNode) {
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

func main() {
	var no1 *ListNode
	//no1 = AddListNode(8,no1)
	//no1 = AddListNode(1,no1)
	//P(no1)

	var no2 *ListNode
	no2 = AddListNode(6,no2)
	//P(no2)

	no := mergeTwoLists(no1,no2)
	P(no)


	//list1 := List{}
	//list1.Add(8)
	//list1.Add(4)
	//list1.Add(1)
	//
	//list2 := List{}
	//list2.Add(8)
	//list2.Add(6)
	//list2.Add(3)
	//list2.Add(1)
	//
	//no := mergeTwoLists2(list1,list2)
	//P(no)
}