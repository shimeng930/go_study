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
		// [2,1],[2,2],[2],[1,1],[4,1],[2]
		lru := ConstructorLRUCache(2)
		lru.Put(1, 1)
		lru.Put(2, 2)
		lru.Get(2)
		lru.Put(1, 1)
		lru.Put(4, 1)
		lru.Get(2)

		//lru.Put(2, 2)
		//lru.Get(1)
		//lru.Put(3, 3)
		//lru.Get(2)
		//lru.Put(4, 4)
		//lru.Get(1)
		//lru.Get(3)
		//lru.Get(4)

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
		var preSum = map[int]int{0: 1}
		preSum[2] = 1
		preSum[2]++

		l := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
		reverseKGroupV1(l, 2)
		reverseKGroup(l, 3)

	})
	t.Run("SHeap", func(t *testing.T) {
		l := &Node{Val: 7}
		l1 := &Node{Val: 13}
		l2 := &Node{Val: 11}
		l3 := &Node{Val: 10}
		l4 := &Node{Val: 1}
		l.Next = l1
		l1.Next = l2
		l1.Random = l
		l2.Next = l3
		l2.Random = l4
		l3.Next = l4
		l3.Random = l2
		l4.Random = l
		copyRandomList(l)

	})
}
