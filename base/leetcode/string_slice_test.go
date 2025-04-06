package leetcode

import (
	"fmt"
	"testing"
)

func Test_longestConsecutive(t *testing.T) {
	t.Run("ongestConsecutive", func(t *testing.T) {
		subarraySum([]int{1, 1, 1}, 2)

		longestConsecutive([]int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6})
	})
}

func Test_maxSlidingWindow(t *testing.T) {
	t.Run("maxSlidingWindow", func(t *testing.T) {
		res := maxSlidingWindow([]int{1, 3, 1, 2, 0, 5}, 3)
		fmt.Println(res)
	})
}

func Test_SHeap(t *testing.T) {
	t.Run("SHeap", func(t *testing.T) {
		res := NewSHeap([]int{0, 99, 5, 36, 7, 22, 17, 46, 12, 2, 19, 25, 28, 1, 92}).BuildHeap()
		fmt.Println(res)

		findKthLargest([]int{-1, 2, 0}, 2)
	})
}

func Test_binarySearch(t *testing.T) {
	t.Run("binarySearch", func(t *testing.T) {
		fmt.Println(binarySearch([]int{1, 3, 5, 6}, 4))
		fmt.Println(searchInsert([]int{3, 5, 7, 9, 10}, 8))
	})
}

func Test_greedy(t *testing.T) {
	t.Run("partitionLabels", func(t *testing.T) {
		fmt.Println(partitionLabels("eaaaabaaec"))
		fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
	})
	t.Run("numDecodings", func(t *testing.T) {
		fmt.Println(numDecodings("12"))
	})
}

func Test_minWindow(t *testing.T) {
	t.Run("partitionLabels", func(t *testing.T) {
		fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	})
	t.Run("checkInclusion", func(t *testing.T) {
		fmt.Println(lengthOfLongestSubstring("abba"))
		fmt.Println(checkInclusion("ab", "eidbaooo"))
	})
}
