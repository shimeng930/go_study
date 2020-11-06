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

func (l *ListNode) GetLastKNode(index int, head *ListNode) {
	if head == nil {
		return
	}

	var slowPtr = head
	var fastPtr = head
	for i:=0; i<index-1; i++ {
		fastPtr = fastPtr.Next
		if fastPtr == nil {
			fmt.Printf("node length(%d) is less than index(%d)", i+1, index)
		}
	}

	for fastPtr.Next != nil {
		fastPtr = fastPtr.Next
		slowPtr = slowPtr.Next
	}
	fmt.Printf("last %d node is %d", index, slowPtr.Val)
}