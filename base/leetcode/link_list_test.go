package leetcode

import (
	"fmt"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	t.Run("SHeap", func(t *testing.T) {
		l := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
		reorderList(l)

		l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
		l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 5}}}
		addTwoNumbers(l1, l2)
	})
}

func Test_LRUCache(t *testing.T) {
	t.Run("LRUCache", func(t *testing.T) {
		lru := ConstructorLRUCache(3)
		//lru.Put(2, 1)
		//lru.Get(2)
		//lru.Put(3, 2)
		//lru.Get(2)
		//lru.Get(3)

		lru.Put(1, 1)
		lru.Put(2, 2)
		lru.Put(3, 3)
		lru.Put(4, 4)
		fmt.Println("lru.Get(4) is ", lru.Get(4))
		fmt.Println("lru.Get(3) is ", lru.Get(3))
		fmt.Println("lru.Get(2) is ", lru.Get(2))
		fmt.Println("lru.Get(1) is ", lru.Get(1))
		lru.Put(5, 5)
		fmt.Println("lru.Get(1) is ", lru.Get(1))
		fmt.Println("lru.Get(2) is ", lru.Get(2))
		fmt.Println("lru.Get(3) is ", lru.Get(3))
		fmt.Println("lru.Get(4) is ", lru.Get(4))
		fmt.Println("lru.Get(5) is ", lru.Get(5))
	})
}

func Test_reverseKGroup(t *testing.T) {
	t.Run("SHeap", func(t *testing.T) {
		l := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
		reverseKGroup(l, 3)

	})
}
