package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var h, cur *ListNode
	var preCar int
	for l1 != nil || l2 != nil || preCar > 0 {
		if l1 == nil {
			l1 = new(ListNode)
		}
		if l2 == nil {
			l2 = new(ListNode)
		}
		sum := l1.Val + l2.Val + preCar
		car, val := sum/10, sum%10
		if h == nil {
			h = &ListNode{Val: val}
			cur = h
		} else {
			cur.Next = &ListNode{Val: val}
			cur = cur.Next
		}
		preCar = car
		l1 = l1.Next
		l2 = l2.Next
	}
	return h
}

// lc-21
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list2.Next, list1)
		return list2
	}
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

// lc-146 LRU
type linkNode struct {
	pre      *linkNode
	next     *linkNode
	key, val int
}

type LRUCache struct {
	capacity int
	used     int
	valueMap map[int]*linkNode
	latest   *linkNode
	tail     *linkNode
}

func ConstructorLRUCache(capacity int) LRUCache {
	return LRUCache{capacity: capacity, valueMap: make(map[int]*linkNode)}
}

func (this *LRUCache) Get(key int) int {
	n := this.valueMap[key]
	if n != nil {
		this.use(n)
		return n.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	n := this.valueMap[key]
	if n == nil {
		node := &linkNode{key: key, val: value}
		this.valueMap[key] = node

		if this.used == this.capacity {
			// del the tail
			this.remove(this.tail.key)
		}
		this.add(node)
	} else {
		n.val = value
		this.use(n)
	}
}

func (this *LRUCache) use(n *linkNode) {
	if n == this.latest {
		return
	}
	if n == this.tail {
		this.tail = this.tail.pre
	} else {
		n.pre.next = n.next
		n.next.pre = n.pre
	}

	//n.pre.next

	n.next = this.latest
	this.latest.pre = n
	this.latest = n
	n.pre = nil
}

func (this *LRUCache) add(node *linkNode) {
	// first time
	this.used++
	if this.latest == nil {
		this.latest = node
		this.tail = node
		return
	}

	node.next = this.latest
	this.latest.pre = node
	this.latest = node
}

func (this *LRUCache) remove(key int) {
	node := this.valueMap[key]
	this.used--
	this.valueMap[key] = nil
	if node == this.latest {
		this.latest, this.tail = nil, nil
		return
	}
	if node == this.tail {
		this.tail = node.pre
		node.pre.next = nil
	}
	// not happen
}

// lc 206 reverse list
func reverseList(head *ListNode) *ListNode {
	// var prev *ListNode
	// for head != nil {
	//     tmp := head.Next
	//     head.Next = prev
	//     prev = head
	//     head = tmp
	// }
	// return prev

	dum := new(ListNode)
	rev(head, dum)
	return dum.Next
}

func rev(head, dum *ListNode) *ListNode {
	if head.Next == nil {
		dum.Next = head
		return head
	}
	h := rev(head.Next, dum)
	h.Next = head
	head.Next = nil
	return head
}

// lc 143  重排链表
func reorderList(head *ListNode) {
	mid := findMid(head)
	l2 := mid.Next
	mid.Next = nil
	l2 = reverse(l2)

	var cur = head
	for cur != nil {
		if l2 != nil {
			tmp := cur.Next
			cur.Next = l2
			l2 = l2.Next
			cur.Next.Next = tmp
			cur = tmp
		} else {
			break
		}
	}
}

func findMid(head *ListNode) *ListNode {
	var f, l = head, head
	for f != nil && f.Next != nil {
		f = f.Next.Next
		l = l.Next
	}
	return l
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = prev
		prev = head
		head = tmp
	}
	return prev
}

// lc 25 n个一组反转链表
func reverseKGroup(head *ListNode, k int) *ListNode {
	f, last, next := reverseK(head, k)
	var nf, nl *ListNode
	for next != nil {
		nf, nl, next = reverseK(next, k)
		last.Next = nf
		if nl != nil {
			last = nl
		}
	}
	return f
}

func reverseK(h *ListNode, k int) (*ListNode, *ListNode, *ListNode) {
	var cur1 = h
	for i := 0; i < k; i++ {
		if cur1 != nil {
			cur1 = cur1.Next
		} else if i < k {
			return h, nil, nil
		}
	}

	var pre *ListNode
	var cur = h
	for cur != nil && k > 0 {
		tmp := cur.Next
		cur.Next = pre
		k--
		pre = cur
		cur = tmp
	}
	if k > 0 {
		return h, nil, nil
	}
	// h.Next = cur
	return pre, h, cur
}
