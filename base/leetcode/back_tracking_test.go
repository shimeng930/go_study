package leetcode

import (
	"fmt"
	"testing"
)

func Test_back_tracking(t *testing.T) {
	t.Run("letterCombinations", func(t *testing.T) {
		fmt.Println(letterCombinations("23"))
	})

	t.Run("listAll", func(t *testing.T) {
		fmt.Println(listAll([]int{1, 2, 3}))
	})
	t.Run("listSub", func(t *testing.T) {
		fmt.Println(listSub([]int{1, 2, 3}, -1))
		fmt.Println(listSub([]int{1, 2, 3}, 2))
	})
}

func Test_bfs(t *testing.T) {
	t.Run("bfs", func(t *testing.T) {
		root := &TreeNode{Val: 3}
		root.Left = &TreeNode{Val: 9}
		root.Right = &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}
		bfsTree2(root)
		bfsTree(root)
	})

}
