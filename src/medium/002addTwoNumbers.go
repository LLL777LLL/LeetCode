package main

import (
	"fmt"
)

/*给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 
一位 数字。如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
示例：
输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/

type Node struct {
	Data int
	Next *Node
}

type List struct {
	headNode *Node
}

//尾插法
func (list *List) Append(data int) {
	node := &Node{Data:data}
	if list.headNode == nil {
		list.headNode = node
	}else {
		cur := list.headNode
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = node
	}
}

//输出 单链表  （为了打印输出之后方便看结果，就将单链表的数据放入了切片）
func (list *List) printList() {
	var nums []int
	pre := list.headNode
	for pre != nil {
		nums = append(nums,pre.Data)
		pre = pre.Next
	}
	fmt.Println(nums)
}

//两数之和
func AddTwoNumbers(l1,l2 List) *Node {
	var node *Node
	flag := 0
	for l1.headNode != nil || l2.headNode != nil || flag != 0 {
		tmp := 0
		if l1.headNode != nil {
			tmp += l1.headNode.Data
			l1.headNode = l1.headNode.Next
		}
		if l2.headNode != nil {
			tmp += l2.headNode.Data
			l2.headNode = l2.headNode.Next
		}
		tmp += flag

		flag = tmp / 10
		tmp %= 10

		no := &Node{Data:tmp}
		if node == nil {
			node = no
		}else {
			cur := node
			for cur.Next != nil {
				cur = cur.Next
			}
			cur.Next = no
		}
	}
	return node
}

func main() {
	list1 := List{}
	list1.Append(2)
	list1.Append(4)
	list1.Append(3)
	list1.printList()

	list2 := List{}
	list2.Append(5)
	list2.Append(6)
	list2.Append(4)
	list2.printList()

	//这里参数是 list 是为了更好的过去两条单链表数据
	node := AddTwoNumbers(list1,list2)

	for node != nil {
		fmt.Println(node.Data)
		node = node.Next
	}
}