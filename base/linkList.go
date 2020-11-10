package base

import (
	"fmt"
	"math/rand"
	"time"
)

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
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=1; i<length; i++ {
		node := &ListNode{Val: r.Intn(50)}
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
			fmt.Printf("node length(%d) is less than index(%d)\n", i+1, index)
			return
		}
	}

	for fastPtr.Next != nil {
		fastPtr = fastPtr.Next
		slowPtr = slowPtr.Next
	}
	fmt.Printf("last %d node is %d\n", index, slowPtr.Val)
}

// 合并有序链表
func (l *ListNode) Merge2List(headA, headB *ListNode) *ListNode {
	if headA == nil {
		return headB
	}
	if headB == nil {
		return headA
	}

	var ptrA = headA
	var ptrB = headB
	var cur *ListNode
	if ptrA.Val < ptrB.Val {
		cur = ptrA
		ptrA = ptrA.Next
	} else {
		cur = ptrB
		ptrB = ptrB.Next
	}
	for ptrA != nil || ptrB != nil {
		if ptrA == nil {
			cur.Next = ptrB
			break
		}
		if ptrB == nil {
			cur.Next = ptrA
			break
		}
		if ptrA.Val < ptrB.Val {
			cur.Next = ptrA
			ptrA = ptrA.Next
		} else {
			cur.Next = ptrB
			ptrB = ptrB.Next
		}
		cur = cur.Next
	}

	if headA.Val < headB.Val {
		return headA
	} else {
		return headB
	}
}

// 效率最低的排序 n*n
func (l *ListNode) SortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var newHead = &ListNode{Val: head.Val}
	var ptr = head.Next
	for ptr != nil {
		pre := newHead
		cur := pre.Next
		node := &ListNode{Val: ptr.Val}
		for pre != nil {
			if cur == nil {
				if ptr.Val > pre.Val {
					pre.Next = node
				} else {
					node.Next = pre
					newHead = node
				}
				break
			}

			if node.Val > pre.Val {
				if node.Val > cur.Val {
					pre = cur
					cur = cur.Next
				} else {
					node.Next = cur
					pre.Next = node
					break
				}
			} else {
				node.Next = pre
				newHead = node
				break
			}
		}
		ptr = ptr.Next
	}
	return newHead
}