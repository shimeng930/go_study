package leetcode

import (
	"fmt"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	t.Run("inorder", func(t *testing.T) {
		root := &TreeNode{Val: 1}
		root.Right = &TreeNode{Val: 2}
		root.Right.Left = &TreeNode{Val: 3}

		inorderTraversal(root)

	})
}

func Test_flatten(t *testing.T) {
	t.Run("flatten", func(t *testing.T) {
		root := &TreeNode{Val: 1}
		root.Left = &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}
		root.Right = &TreeNode{Val: 5, Right: &TreeNode{Val: 6}}

		//flatten(root)
		//flattenV2(root)
		longestConsecutive([]int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6})
	})
}

func Test_sortedArrayToBST(t *testing.T) {
	t.Run("sortedArrayToBST", func(t *testing.T) {
		arr := []int{2, 0, 33, -1, 1, 25, 40, -1, -1, 11, 31, 34, 45, 10, 18, 29, 32, -1, 36, 43, 46, 4, -1, 12, 24, 26, 30, -1, -1, 35, 39, 42, 44,
			-1, 48, 3, 9, -1, 14, 22, -1, -1, 27, -1, -1, -1, -1, 38, -1, 41, -1, -1, -1, 47, 49, -1, -1, 5, -1, 13, 15, 21, 23, -1, 28, 37, -1, -1, -1, -1,
			-1, -1, -1, -1, 8, -1, -1, -1, 17, 19, -1, -1, -1, -1, -1, -1, -1, 7, -1, 16, -1, -1, 20, 6}
		r := convertTree(arr)

		deleteNode(r, 33)

		root := sortedArrayToBST([]int{-10, -3, 0, 5, 9})
		fmt.Println(root)
	})
}

func Test_buildTree(t *testing.T) {
	t.Run("sortedArrayToBST", func(t *testing.T) {

		fmt.Println(findKid([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
	})
	t.Run("buildTreePI", func(t *testing.T) {
		r := buildTreePI([]int{3,9,20,15,7}, []int{9,3,15,20,7})
		fmt.Println(r)
	})
	t.Run("diameterOfBinaryTree", func(t *testing.T) {
		root := &TreeNode{Val: 1}
		root.Left = &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}
		root.Right = &TreeNode{Val: 5}
		res := diameterOfBinaryTree(root)
		fmt.Println(res)
	})
	t.Run("flattenV", func(t *testing.T) {
		root := &TreeNode{Val: 1}
		root.Left = &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}
		root.Right = &TreeNode{Val: 5, Right: &TreeNode{Val: 6}}
		flattenV(root)
		fmt.Println(root)
	})
	t.Run("maxDepth", func(t *testing.T) {
		root := &TreeNode{Val: 1}
		root.Left = &TreeNode{Val: 2}
		root.Right = &TreeNode{Val: 3, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}
		fmt.Println(maxDepth(root))
	})
}

func addLeft(node *TreeNode, n int) *TreeNode {
	if node == nil || n == -1 {
		return nil
	}
	node.Left = &TreeNode{Val: n}
	return node.Left
}
func addRight(node *TreeNode, n int) *TreeNode {
	if node == nil || n == -1 {
		return nil
	}
	node.Right = &TreeNode{Val: n}
	return node.Right
}

func convertTree(arrs []int) *TreeNode {
	if len(arrs) == 0 {
		return nil
	}

	//var level = []*TreeNode{{Val: arrs[0]}}

	var buildTreeHelper func(arr []int, level []*TreeNode)
	buildTreeHelper = func(arr []int, level []*TreeNode) {
		if len(arr) == 0 {
			return
		}
		var kid int
		var newLevel []*TreeNode
		for _, item := range level {
			if arr[kid] != -1 {
				l := &TreeNode{Val: arr[kid]}
				newLevel = append(newLevel, l)
				item.Left = l
			}
			kid++
			if kid >= len(arr) {
				break
			}
			if arr[kid] != -1 {
				r := &TreeNode{Val: arr[kid]}
				newLevel = append(newLevel, r)
				item.Right = r
			}
			kid++
			if kid >= len(arr) {
				break
			}
		}
		if kid >= len(arr) {
			return
		}
		buildTreeHelper(arr[kid:], newLevel)

		return
	}
	root := &TreeNode{Val: arrs[0]}
	buildTreeHelper(arrs[1:], []*TreeNode{root})
	return root
}
