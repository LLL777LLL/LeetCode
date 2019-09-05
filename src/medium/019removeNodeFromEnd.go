package main

import "fmt"

/*给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
示例：
给定一个链表: 1->2->3->4->5, 和 n = 2.
当删除了倒数第二个节点后，链表变为 1->2->3->5.

说明：给定的 n 保证是有效的。
进阶：你能尝试使用一趟扫描实现吗？

*/

type ListNode struct {
	Val int
	Next *ListNode
}

//头插法
func Add(n int,li *ListNode) *ListNode {
	no := &ListNode{Val:n}
	no.Next = li
	return no
}

//尾插法
func Append(n int,li *ListNode) *ListNode {
	node := &ListNode{Val:n}
	if li == nil {
		li = node
	}else {
		no := li
		for no.Next != nil {
			no = no.Next
		}
		no.Next = node
	}
	return li
}


func printList(li *ListNode) {
	var arr []int
	for li != nil {
		arr = append(arr,li.Val)
		li = li.Next
	}
	fmt.Println(arr)
}

//一次遍历
// ---- 上述算法可以优化为只使用一次遍历。我们可以使用两个指针而不是一个指针。第一个指针从列表的开头向前移动 n+1n+1 步，而第二个指针将从列表的开头出发。
// 现在，这两个指针被 nn 个结点分开。我们通过同时移动两个指针向前来保持这个恒定的间隔，直到第一个指针到达最后一个结点。此时第二个指针将指向从最后一个结
// 点数起的第 nn 个结点。我们重新链接第二个指针所引用的结点的 next 指针指向该结点的下下个结点。
func removeNodeFromEnd(head *ListNode,n int) *ListNode {
	last := &ListNode{}
	last.Next = head   //考虑到删除倒数最后一个，也就是删除第一个。这样的话在删除不是倒数第一个的时候，就要用 last.Next != nil 来判断，因为这里加了一个头节点
	for i := 0;i < n+1;i++ {
		last = last.Next
	}
	first := head
	if last == nil {   //删除倒数第一个时的情况
		first = head.Next
		return first
	}
	for last.Next != nil {
		last = last.Next
		first = first.Next
	}
	first.Next = first.Next.Next
	return head
}

//常规思想（两次遍历 -- 先算出单链表长度，之后再删除）
func removeNodeFromEnd2(li *ListNode,n int) *ListNode {
	l := 0
	l1 := li
	for l1 != nil {
		l++
		l1 = l1.Next
	}
	printList(li)

	p := li
	for i := 0;i < l-n-1;i++ {
		p = p.Next
	}
	p.Next = p.Next.Next
	return li
}

func main() {
	//头插法
	var l1 *ListNode
	l1 = Add(6,l1)
	l1 = Add(5,l1)
	l1 = Add(4,l1)
	l1 = Add(3,l1)
	l1 = Add(2,l1)
	l1 = Add(1,l1)
	printList(l1)

	l11 := removeNodeFromEnd(l1,1)
	printList(l11)

	//尾插法
	//var l2 *ListNode
	//l2 = Append(1,l2)
	//l2 = Append(2,l2)
	//l2 = Append(3,l2)
	//printList(l2)
}