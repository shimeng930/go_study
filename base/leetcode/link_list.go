package leetcode

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) print() {
	fmt.Print("list value are: ")
	var head = l
	for head != nil {
		fmt.Print(head.Val)
		head = head.Next
		if head != nil {
			fmt.Print("->")
		}
	}
	fmt.Println()
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var dum = new(ListNode)
	dum.Next = head
	pre := findEndK(dum, n+1)
	pre.Next = pre.Next.Next
	return dum.Next
}

func findEndK(h *ListNode, k int) *ListNode {
	var p1 = h
	for i := 0; i < k; i++ {
		p1 = p1.Next
	}
	var p2 = h
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p2
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

// lc-21 合并两个有序链表
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

// lc 21 合并有序链表
func (l *ListNode) merge2List(list1, list2 *ListNode) *ListNode {
	var n = new(ListNode)
	var cur = n
	for list1 != nil || list2 != nil {
		if list1 == nil {
			n.Next = list2
			break
		}
		if list2 == nil {
			n.Next = list1
			break
		}

		if list1.Val < list2.Val {
			n.Next = list1
			list1 = list1.Next
		} else {
			n.Next = list2
			list2 = list2.Next
		}
		n = n.Next
	}
	return cur.Next
}

// lc-146 LRU
//type linkNode struct {
//	pre      *linkNode
//	next     *linkNode
//	key, val int
//}
//
//type LRUCache struct {
//	capacity int
//	used     int
//	valueMap map[int]*linkNode
//	latest   *linkNode
//	tail     *linkNode
//}
//
//func ConstructorLRUCache(capacity int) LRUCache {
//	return LRUCache{capacity: capacity, valueMap: make(map[int]*linkNode)}
//}
//
//func (this *LRUCache) Get(key int) int {
//	n := this.valueMap[key]
//	if n != nil {
//		this.use(n)
//		return n.val
//	}
//	return -1
//}
//
//func (this *LRUCache) Put(key int, value int) {
//	n := this.valueMap[key]
//	if n == nil {
//		node := &linkNode{key: key, val: value}
//		this.valueMap[key] = node
//
//		if this.used == this.capacity {
//			// del the tail
//			this.remove(this.tail.key)
//		}
//		this.add(node)
//	} else {
//		n.val = value
//		this.use(n)
//	}
//}
//
//func (this *LRUCache) use(n *linkNode) {
//	if n == this.latest {
//		return
//	}
//	if n == this.tail {
//		this.tail = this.tail.pre
//	} else {
//		n.pre.next = n.next
//		n.next.pre = n.pre
//	}
//
//	//n.pre.next
//
//	n.next = this.latest
//	this.latest.pre = n
//	this.latest = n
//	n.pre = nil
//}
//
//func (this *LRUCache) add(node *linkNode) {
//	// first time
//	this.used++
//	if this.latest == nil {
//		this.latest = node
//		this.tail = node
//		return
//	}
//
//	node.next = this.latest
//	this.latest.pre = node
//	this.latest = node
//}
//
//func (this *LRUCache) remove(key int) {
//	node := this.valueMap[key]
//	this.used--
//	this.valueMap[key] = nil
//	if node == this.latest {
//		this.latest, this.tail = nil, nil
//		return
//	}
//	if node == this.tail {
//		this.tail = node.pre
//		node.pre.next = nil
//	}
//	// not happen
//}

// lc 206 reverseList
func reverseList(head *ListNode) *ListNode {
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

// lc.24 两两交换链表中的节点
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var next = head.Next
	head.Next = swapPairs(next.Next)
	next.Next = head
	return next
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

func reverseKGroupV1(head *ListNode, k int) *ListNode {
	var size int
	n := head
	for n != nil {
		size += 1
		n = n.Next
	}

	var dum = new(ListNode)
	dum.Next = head
	var pre, cur *ListNode = nil, dum.Next
	var temp = dum
	for size >= k {
		size -= k
		for i := 0; i < k; i++ {
			tmp := cur.Next
			cur.Next = pre
			pre = cur
			cur = tmp
		}
		tmp := temp.Next
		temp.Next.Next = cur
		temp.Next = pre
		temp = tmp
	}
	return dum.Next
}

func reverseKGroupV2(head *ListNode, k int) *ListNode {
	var ng = head
	for i := 0; i < k; i++ {
		if ng == nil {
			return head
		}
		ng = ng.Next
	}
	var tmp = head
	// 这里[tmp, ng) 代表需要反转的一组，左闭右开。
	// 返回值是反转之后的头节点
	// 然后ng作为下一组的头结点 进行递归
	rh := reverseBeforeTail(tmp, ng)
	tmp.Next = reverseKGroupV2(ng, k)
	return rh
}

func reverseBeforeTail(h, t *ListNode) *ListNode {
	var pre *ListNode
	var cur = h
	for cur != t {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

//type LinkNode struct {
//	Key, Val  int
//	Pre, Next *LinkNode
//}
//
//type LRUCache struct {
//	Capacity, Used int
//	mapIndex       map[int]*LinkNode // data key : node
//	head, tail     *LinkNode
//}
//
//func ConstructorLRUCache(capacity int) LRUCache {
//	return LRUCache{
//		Capacity: capacity,
//		head:     new(LinkNode),
//		mapIndex: make(map[int]*LinkNode),
//	}
//}
//
//func (l *LRUCache) Get(key int) int {
//	if l.mapIndex[key] != nil {
//		l.adjust(l.mapIndex[key])
//		return l.mapIndex[key].Val
//	}
//	return -1
//}
//
//func (l *LRUCache) adjust(node *LinkNode) {
//	if l.Used == 1 {
//		return
//	}
//	if node == l.tail {
//		l.changeTail()
//	} else {
//		node.Pre.Next = node.Next
//		node.Next.Pre = node.Pre
//	}
//	l.toHead(node)
//}
//func (l *LRUCache) toHead(node *LinkNode) {
//	if node == l.head.Next {
//		return
//	}
//
//	tmp := l.head.Next
//	tmp.Pre = node
//	node.Next = tmp
//	node.Pre = l.head
//	l.head.Next = node
//}
//func (l *LRUCache) changeTail() {
//	l.tail = l.tail.Pre
//	l.tail.Next = nil
//}
//
//func (l *LRUCache) Put(key int, value int) {
//	var n = l.mapIndex[key]
//	if n != nil {
//		n.Val = value
//		l.adjust(n)
//		return
//	}
//	if l.Used == l.Capacity {
//		delete(l.mapIndex, l.tail.Key)
//		l.changeTail()
//		l.Used--
//	}
//	var nn = &LinkNode{Key: key, Val: value}
//	l.mapIndex[key] = nn
//	l.Used++
//	if l.Used == 1 {
//		l.head.Next = nn
//		nn.Pre = l.head
//		l.tail = nn
//	}
//	l.toHead(nn)
//}

type LRUCache struct {
	capacity   int
	used       int
	value      map[int]int
	head, tail *LinkNode
}

type LinkNode struct {
	Val       int
	Pre, Next *LinkNode
}

func ConstructorLRU(capacity int) LRUCache {
	v := make(map[int]int)
	c := LRUCache{value: v, capacity: capacity}
	c.head = new(LinkNode)
	c.tail = new(LinkNode)
	c.head.Next = c.tail
	c.tail.Pre = c.head
	return c
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.value[key]; ok {
		this.adjust(key, false)
		return this.value[key]
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	var addNew bool
	if _, ok := this.value[key]; !ok {
		addNew = true
		this.used++
	}
	this.value[key] = value
	// adjust
	this.adjust(key, addNew)
}

func (this *LRUCache) adjust(key int, addNew bool) {
	var node = &LinkNode{Val: key}
	node.Next = this.head.Next
	this.head.Next.Pre = node
	node.Pre = this.head
	this.head.Next = node
	if addNew {
		if this.used > this.capacity {
			del := this.tail.Pre
			del.Pre.Next = this.tail
			this.tail.Pre = del.Pre
			this.used--
			delete(this.value, del.Val)
		}
	} else {
		var cur = this.head.Next.Next
		for cur != nil {
			if cur.Val == key {
				cur.Pre.Next = cur.Next
				cur.Next.Pre = cur.Pre
				break
			}
			cur = cur.Next
		}
	}
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	var nodeMap = make(map[*Node][2]*Node)
	var cur = head
	for cur != nil {
		nodeMap[cur] = [2]*Node{cur.Next, cur.Random}
		cur = cur.Next
	}
	var dummy = new(Node)
	dummy.Next = copyN(head, nodeMap)

	return dummy.Next
}

func copyN(node *Node, nm map[*Node][2]*Node) *Node {
	if node == nil {
		return nil
	}
	cp := &Node{
		Val:  node.Val,
		Next: copyN(nm[node][0], nm),
	}
	return cp
}

// lc.23 合并K个链表
// 思路：遍历链表，进行相邻合并，时间复杂度是O(K^2/2 * N)
// 优化：改成归并，两两合并，时间复杂度是O(K*logK * N)
func mergeKLists(lists []*ListNode) *ListNode {
	var pre *ListNode
	for i, n := range lists {
		if i == 0 {
			pre = n
			continue
		}
		pre = merge2(pre, n)
	}
	return pre
}

func merge2(l1, l2 *ListNode) *ListNode {
	var cursor = new(ListNode)
	var node = cursor
	for l1 != nil || l2 != nil {
		if l1 == nil {
			cursor.Next = l2
			break
		}
		if l2 == nil {
			cursor.Next = l1
			break
		}
		if l1.Val < l2.Val {
			cursor.Next = l1
			l1 = l1.Next
		} else {
			cursor.Next = l2
			l2 = l2.Next
		}
		cursor = cursor.Next
	}
	return node.Next
}
