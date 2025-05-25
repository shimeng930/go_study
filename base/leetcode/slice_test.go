package leetcode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_slice(t *testing.T) {
	t.Run("threeSum", func(t *testing.T) {
		fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	})
	t.Run("maxOperations", func(t *testing.T) {
		fmt.Println(maxOperations([]int{3, 1, 2, 5, 5}))
		fmt.Println(maxOperations([]int{4, 2, 5, 5, 4}))
	})
	t.Run("rotate", func(t *testing.T) {
		var nums = []int{-1, -100, 3, 99}
		rotate(nums, 2)
		fmt.Println(nums)
	})
	t.Run("lengthOfLIS", func(t *testing.T) {
		//lengthOfLIS([]int{10,9,2,5,3,7,101,18})
		fmt.Println(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
		fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
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

	t.Run("findMedianSortedArrays", func(t *testing.T) {
		fmt.Println(findMedianSortedArraysV2([]int{2, 3, 4, 8, 9}, []int{1, 3, 5, 6}))
		fmt.Println(findMedianSortedArraysV2([]int{4, 5, 6, 8, 9}, []int{}))
		fmt.Println(findMedianSortedArrays([]int{2, 3, 4, 8, 9}, []int{1, 3, 5, 6}))
		fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
		fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2, 4}))
	})
	t.Run("findAnagrams", func(t *testing.T) {
		fmt.Println(findAnagrams("abab", "ab"))
		fmt.Println(findAnagrams("cbaebabacd", "abc"))
	})

}

func Test_binarySearch(t *testing.T) {
	t.Run("findTarget", func(t *testing.T) {
		fmt.Println(findTarget([]int{8, 9, 3, 4, 5, 6, 7}, 3))
	})
	t.Run("binarySearch", func(t *testing.T) {
		fmt.Println(binarySearch([]int{1, 3, 5, 6}, 4))
		fmt.Println(searchInsert([]int{3, 5, 7, 9, 10}, 8))
	})
}

func Test_maxSlidingWindow(t *testing.T) {
	t.Run("maxSlidingWindow", func(t *testing.T) {
		fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6}, 3))
		fmt.Println(maxSlidingWindow([]int{1, -1}, 1))
		fmt.Println(maxSlidingWindow([]int{1}, 1))
	})
}

func Test_SHeap(t *testing.T) {
	t.Run("SHeap", func(t *testing.T) {
		//res := NewSHeap([]int{0, 99, 5, 36, 7, 22, 17, 46, 12, 2, 19, 25, 28, 1, 92}).BuildHeap()
		//fmt.Println(res)

		findKthLargest([]int{-1, 2, 0}, 2)
	})
}

func Test_multi_slice(t *testing.T) {
	t.Run("spiralOrder", func(t *testing.T) {
		var nums = [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
		fmt.Println(spiralOrder(nums))
	})

	t.Run("setZeroes", func(t *testing.T) {
		var nums = [][]int{{1, 2, 3, 4}, {5, 0, 7, 8}, {0, 10, 11, 12}, {13, 14, 15, 0}}
		setZeroes(nums)
		fmt.Println(nums)
	})

	t.Run("rotateMatrix", func(t *testing.T) {
		var nums = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
		rotateMatrix(nums)
		fmt.Println(nums)
	})

}
