package leetcode

import (
	"fmt"
	"math"
	"testing"
)

const nilNode = math.MinInt32

func Test_inorderTraversal(t *testing.T) {
	t.Run("inorder", func(t *testing.T) {
		root := &TreeNode{Val: 1}
		root.Right = &TreeNode{Val: 2}
		root.Right.Left = &TreeNode{Val: 3}

		inorderTraversal(root)
	})
	t.Run("levelOrder", func(t *testing.T) {
		root := &TreeNode{Val: 3}
		root.Left = &TreeNode{Val: 9}
		root.Right = &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}
		fmt.Println(levelOrder(root))
	})
	t.Run("isValidBST", func(t *testing.T) {
		root := &TreeNode{Val: 5}
		root.Left = &TreeNode{Val: 4}
		root.Right = &TreeNode{Val: 6}
		root.Right.Left = &TreeNode{Val: 3}
		root.Right.Right = &TreeNode{Val: 7}

		isValidBST(root)
	})
	t.Run("kthSmall", func(t *testing.T) {
		root := &TreeNode{Val: 3}
		root.Left = &TreeNode{Val: 1}
		root.Right = &TreeNode{Val: 4}
		root.Left.Right = &TreeNode{Val: 2}

		kthSmall(root, 1)
	})
	t.Run("pathSum", func(t *testing.T) {
		var arr = []int{10, 5, -3, 3, 2, nilNode, 11, 3, -2, nilNode, 1}
		arr = []int{1, -2, -3, 1, 3, -2, nilNode, -1}
		root := convertTree(arr)
		pathSum(root, -1)
		//pathSum(root, 8)
	})
	t.Run("pathSumV1", func(t *testing.T) {
		root := &TreeNode{Val: 1}
		root.Left = &TreeNode{Val: -2}
		root.Right = &TreeNode{Val: -3}
		fmt.Println(pathSumV1(root, -1))
	})
	t.Run("lowestCommonAncestorV1", func(t *testing.T) {
		root := &TreeNode{Val: 1}
		root.Left = &TreeNode{Val: 2, Right: &TreeNode{Val: 4}}
		root.Right = &TreeNode{Val: 3}
		fmt.Println(lowestCommonAncestorV1(root, &TreeNode{Val: 4}, &TreeNode{Val: 3}))
	})
}

func Test_flatten(t *testing.T) {
	t.Run("flatten", func(t *testing.T) {
		root := &TreeNode{Val: 1}
		root.Left = &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}
		root.Right = &TreeNode{Val: 5, Right: &TreeNode{Val: 6}}
		treeLengthSum(root)

		//flatten(root)
		//flattenV2(root)
	})
	t.Run("flattenV", func(t *testing.T) {
		root := &TreeNode{Val: 1}
		root.Left = &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}
		root.Right = &TreeNode{Val: 5, Right: &TreeNode{Val: 6}}
		flattenV(root)
		fmt.Println(root)
	})
}

func Test_sortedArrayToBST(t *testing.T) {
	t.Run("sortedArrayToBST", func(t *testing.T) {
		arr := []int{2, 0, 33, nilNode, 1, 25, 40, nilNode, nilNode, 11, 31, 34, 45, 10, 18, 29, 32, nilNode, 36, 43, 46, 4, nilNode, 12, 24, 26, 30, nilNode, nilNode, 35, 39, 42, 44,
			nilNode, 48, 3, 9, nilNode, 14, 22, nilNode, nilNode, 27, nilNode, nilNode, nilNode, nilNode, 38, nilNode, 41, nilNode, nilNode, nilNode, 47, 49, nilNode, nilNode, 5, nilNode, 13, 15, 21, 23, nilNode, 28, 37, nilNode, nilNode, nilNode, nilNode,
			nilNode, nilNode, nilNode, nilNode, 8, nilNode, nilNode, nilNode, 17, 19, nilNode, nilNode, nilNode, nilNode, nilNode, nilNode, nilNode, 7, nilNode, 16, nilNode, nilNode, 20, 6}
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
		r := buildTreePI([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
		fmt.Println(r)
	})
	t.Run("diameterOfBinaryTree", func(t *testing.T) {
		root := &TreeNode{Val: 1}
		root.Left = &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}
		root.Right = &TreeNode{Val: 5}
		res := diameterOfBinaryTree(root)
		fmt.Println(res)
	})
	t.Run("maxDepth", func(t *testing.T) {
		root := &TreeNode{Val: 1}
		root.Left = &TreeNode{Val: 2}
		root.Right = &TreeNode{Val: 3, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}
		fmt.Println(maxDepth(root))
	})
}

func addLeft(node *TreeNode, n int) *TreeNode {
	if node == nil || n == nilNode {
		return nil
	}
	node.Left = &TreeNode{Val: n}
	return node.Left
}
func addRight(node *TreeNode, n int) *TreeNode {
	if node == nil || n == nilNode {
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
			if arr[kid] != nilNode {
				l := &TreeNode{Val: arr[kid]}
				newLevel = append(newLevel, l)
				item.Left = l
			}
			kid++
			if kid >= len(arr) {
				break
			}
			if arr[kid] != nilNode {
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
