package leetcode

import (
	"fmt"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var res []int
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
		return
	}
	inorder(root)
	return res
}

// 二叉树转链表
func flatten(root *TreeNode) {
	var dum []*TreeNode
	var preorder func(*TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		dum = append(dum, node)
		preorder(node.Left)
		preorder(node.Right)
	}

	preorder(root)

	for i := 1; i < len(dum); i++ {
		prev, curr := dum[i-1], dum[i]
		prev.Left, prev.Right = nil, curr
	}
}

// 二叉树转链表
// 左子树move到右子树
func flattenV2(root *TreeNode) {
	//if root == nil {
	//	return
	//}

	var cur = root

	for cur != nil {
		// 左子树为 nil，直接考虑下一个节点
		if cur.Left == nil {
			//root = root.Right
		} else {
			// 找左子树最右边的节点
			pre := cur.Left
			for pre.Right != nil {
				pre = pre.Right
			}
			// 将原来的右子树接到左子树的最右边节点
			pre.Right = cur.Right
			// 将左子树插入到右子树的地方
			cur.Right = cur.Left
			cur.Left = nil
		}
		// 考虑下一个节点
		//root = root.Right
		cur = cur.Right
	}
	fmt.Println(root)
}

// array to bst
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	//var root = new(TreeNode)
	//var toBst func(numsx []int, node *TreeNode)
	//toBst = func(numsx []int, node *TreeNode) {
	//	if len(numsx) == 1 {
	//		node.Val = numsx[0]
	//		return
	//	}
	//
	//
	//	mid := len(numsx)/2
	//	node.Val = numsx[mid]
	//	left := numsx[:mid]
	//	right := numsx[mid+1:]
	//	if len(left) > 0 {
	//		node.Left = new(TreeNode)
	//		toBst(numsx[:mid], node.Left)
	//	}
	//	if len(right) > 0 {
	//		node.Right = new(TreeNode)
	//		toBst(numsx[mid+1:], node.Right)
	//	}
	//}
	//
	//toBst(nums, root)
	//if len(nums) == 1 {
	//	return &TreeNode{Val: nums[0]}
	//}
	//
	mid := len(nums) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root
}

// lc-230
func kthSmallest(root *TreeNode, k int) int {
	var target *TreeNode
	var inOrder func(node *TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil || target != nil {
			return
		}
		inOrder(node.Left)
		// nums = append(nums, node.Val)
		k--
		if k == 0 {
			target = node
		}
		inOrder(node.Right)
	}
	inOrder(root)
	return target.Val
}

// SHeap ------------------------------------------------
// 初始化堆的时候，需要把数组0当成虚拟数据；这是为下标计算的时候方便，从1开始即可
type SHeap struct {
	arr []int
}

func NewSHeap(arr []int) *SHeap {
	return &SHeap{arr: arr}
}

func (s *SHeap) size() int {
	return len(s.arr) - 1
}

func (s *SHeap) BuildHeap() *SHeap {
	for i := s.size() / 2; i >= 1; i-- {
		s.shiftDown(i)
	}
	return s
}

func (s *SHeap) swap(i, j int) {
	s.arr[i], s.arr[j] = s.arr[j], s.arr[i]
}

// parent < kid
func (s *SHeap) less(i, j int) bool {
	return s.arr[i] < s.arr[j]
}

// parent > kid
func (s *SHeap) large(i, j int) bool {
	return s.arr[i] > s.arr[j]
}

func (s *SHeap) shiftDown(i int) {
	var flag = true
	var t = i
	for i*2 <= s.size() && flag {
		if s.large(i, i*2) {
			t = i * 2
		}
		if i*2+1 <= s.size() {
			if s.large(t, i*2+1) {
				t = i*2 + 1
			}
		}
		if t == i {
			flag = false
		} else {
			s.swap(i, t)
			i = t
		}
	}
}

func (s *SHeap) push(v int) {
	if v <= s.arr[1] {
		return
	}

	s.arr[1] = v
	s.shiftDown(1)
}

func findKthLargest(nums []int, k int) int {
	karr := []int{0} // 为了下标方便计算，空一个节点出来，实际的根节点从1开始
	karr = append(karr, nums[:k]...)
	sh := NewSHeap(karr).BuildHeap()
	for i := k; i < len(nums); i++ {
		sh.push(nums[i])
	}
	return sh.arr[1]
}

// lc 450, delete node in bst
func deleteNode(root *TreeNode, key int) *TreeNode {

	// search
	cur := root
	//var pn *TreeNode
	pn, kn := search(nil, cur, key)
	if kn == nil {
		return root
	}

	// if kn is root
	var adjustKn = adjustCur(kn)
	if pn == nil {
		return adjustKn
	}

	// kn is right of pn
	if pn.Right == kn {
		pn.Right = adjustKn
	} else {
		pn.Left = adjustKn
	}
	return root
}

func adjustCur(kn *TreeNode) *TreeNode {
	if kn.Left != nil && kn.Right != nil {
		// right deep left kid up
		var deepL = kn.Right.Left
		if deepL == nil {
			kn.Right.Left = kn.Left
			return kn.Right
		}
		var deelP = kn.Right
		for deepL.Left != nil {
			deelP = deepL
			deepL = deepL.Left
		}

		deepL.Left = kn.Left
		deepL.Right = kn.Right
		deelP.Left = nil
		return deepL

	} else if kn.Left != nil {
		// left up
		return kn.Left
	}
	return kn.Right
}

func search(p, cur *TreeNode, key int) (*TreeNode, *TreeNode) {
	for cur != nil {
		if cur.Val == key {
			return p, cur
		}
		p = cur
		if cur.Val > key {
			return search(p, cur.Left, key)
		} else {
			return search(p, cur.Right, key)
		}
	}
	return nil, nil
}

// lc 105 covert tree
func findKid(pre []int, inorder []int) *TreeNode {
	if len(pre) == 0 || len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: pre[0]}
	for i, item := range inorder {
		if item == root.Val {
			root.Left = findKid(pre[1:], inorder[:i])
			if root.Left != nil {
				pre = pre[1:]
			} else {
				return root
			}
			root.Right = findKid(pre, inorder[i:])
			if root.Right != nil {
				pre = pre[1:]
			} else {
				return root
			}
		}
	}
	return root

}

// lc 105
func buildTreePI(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	var rootIdx int
	for i, n := range inorder {
		if n == preorder[0] {
			rootIdx = i
			break
		}
	}
	root.Left = buildTreePI(preorder[1:rootIdx+1], inorder[:rootIdx])
	root.Right = buildTreePI(preorder[rootIdx+1:], inorder[rootIdx+1:])

	return root
}

func diameterOfBinaryTree(root *TreeNode) int {
	var res int
	var maxdepth func(r *TreeNode) int
	maxdepth = func(r *TreeNode) int {
		if r == nil {
			return 0
		}
		left := maxdepth(r.Left)
		right := maxdepth(r.Right)
		res = max(res, left+right)
		return max(left, right) + 1
	}

	maxdepth(root)
	return res
}

func flattenV(root *TreeNode) {
	var dum = new(TreeNode)
	var cur = dum

	var flattenDo func(root *TreeNode)
	flattenDo = func(root *TreeNode) {
		if root == nil {
			return
		}
		cur.Right = &TreeNode{Val: root.Val}
		cur = cur.Right
		flattenDo(root.Left)
		flattenDo(root.Right)
	}

	flattenDo(root)
	root = dum.Right
	fmt.Println(root)
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMax := maxDepth(root.Left)
	rightMax := maxDepth(root.Right)
	return 1 + min(leftMax, rightMax)
}

// lc 102 二叉树层序遍历
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	var arr = []*TreeNode{root}
	res = append(res, []int{root.Val})

	for len(arr) > 0 {
		arr = nrlevel(arr)
		le := level(arr)
		res = append(res, le)
	}
	return res
}

func level(arr []*TreeNode) (res []int) {
	for _, item := range arr {
		res = append(res, item.Val)
	}
	return res
}

func nrlevel(arr []*TreeNode) (res []*TreeNode) {
	for _, item := range arr {
		if item.Left != nil {
			res = append(res, item.Left)
		}
		if item.Right != nil {
			res = append(res, item.Right)
		}
	}
	return res
}

// lc 98 验证二叉搜索树
func isValidBST(root *TreeNode) bool {
	// 这个解法是错误的，因为只比较了根节点与直接子节点的大小关系；
	// 但是实际上根节点是要比左边所有都小，比右边所有都大
	//if root.Left == nil && root.Right == nil {
	//	return true
	//}
	//if root.Left == nil && root.Val < root.Right.Val {
	//	return true
	//}
	//if root.Right == nil && root.Val > root.Left.Val {
	//	return true
	//}
	//if isValidBST(root.Left) && isValidBST(root.Right) {
	//	return true
	//}
	//return false

	//var end func(n *TreeNode) bool
	//end = func(n *TreeNode) bool {
	//	if n == nil {
	//		return true
	//	}
	//
	//	if !end(n.Left) || !end(n.Right) {
	//		return false
	//	}
	//	if (n.Left == nil || n.Left.Val < n.Val) &&
	//		(n.Right == nil || n.Val < n.Right.Val) {
	//		return true
	//	}
	//	return false
	//}
	//return end(root)

	var pre *TreeNode
	var in func(node *TreeNode) bool
	in = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		if !in(node.Left) {
			return false
		}
		if pre != nil && pre.Val >= node.Val {
			return false
		}
		pre = node
		return in(node.Right)
	}
	return in(root)
}

func kthSmall(root *TreeNode, k int) int {
	var t *TreeNode
	var in func(n *TreeNode)
	in = func(n *TreeNode) {
		if t != nil {
			return
		}
		if n == nil {
			return
		}
		in(n.Left)
		k--
		if k == 0 {
			t = n
			return
		}
		in(n.Right)
	}
	in(root)
	return t.Val
}
