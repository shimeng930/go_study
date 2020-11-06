package base

import "fmt"

type ListNode struct {
	Val 	int
	Next 	*ListNode
}

func NewListNode() *ListNode {
	return &ListNode{}
}

func (l *ListNode) InitList(length int) *ListNode {
	head := &ListNode{Val: 0}
	temp := head
	for i:=1; i<length; i++ {
		node := &ListNode{Val: i}
		temp.Next = node
		temp = node
	}
	return head
}

func (l *ListNode) Print(head *ListNode)  {
	if head == nil {
		fmt.Println("nil node list")
		return
	}

	var cur = head
	for {
		if cur != nil {
			fmt.Printf("%d->", cur.Val)
			cur = cur.Next
		} else {
			fmt.Println("nil")
			break
		}
	}
}

func (l *ListNode) Reverse(head *ListNode) *ListNode {
	var cur = head
	var pre *ListNode
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}