package leetcode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_slice(t *testing.T) {
	t.Run("rotate", func(t *testing.T) {
		var nums = []int{-1, -100, 3, 99}
		rotate(nums, 2)
		fmt.Println(nums)
	})
	t.Run("reverseUrl", func(t *testing.T) {
		var s = "www.taobao.com"
		reverseUrl(s)
		fmt.Println(s)
	})
	t.Run("lengthOfLIS", func(t *testing.T) {
		//lengthOfLIS([]int{10,9,2,5,3,7,101,18})
		fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	})
	t.Run("findTarget", func(t *testing.T) {
		fmt.Println(findTarget([]int{8, 9, 3, 4, 5, 6, 7}, 3))
	})
	t.Run("removeElement", func(t *testing.T) {
		fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
	})
	t.Run("removeDuplicates", func(t *testing.T) {
		fmt.Println(removeDuplicates([]int{1, 1, 1, 1}))
	})
	t.Run("findMax", func(t *testing.T) {
		assert.Equal(t, 5, findMax([]int{1, 2, 3, 4, 5}))
		assert.Equal(t, 5, findMax([]int{3, 4, 5, 1, 2}))
		assert.Equal(t, 2, findMax([]int{2, 1}))
		assert.Equal(t, 3, findMax([]int{2, 3, 1}))
		// find min
		assert.Equal(t, 1, findMin([]int{3, 4, 5, 1, 2}))
		assert.Equal(t, 1, findMin([]int{2, 1}))
		assert.Equal(t, 1, findMin([]int{2, 3, 1}))
	})

}

func Test_multi_slice(t *testing.T) {
	t.Run("spiralOrder", func(t *testing.T) {
		var nums = [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
		fmt.Println(spiralOrder(nums))
	})

}
