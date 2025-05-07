package leetcode

import (
	"fmt"
	"testing"
)

func Test_longestConsecutive(t *testing.T) {
	t.Run("longestConsecutive", func(t *testing.T) {
		subarraySum([]int{1, 1, 1}, 2)

		longestConsecutive([]int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6})
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

func Test_string(t *testing.T) {
	t.Run("reverseUrl", func(t *testing.T) {
		var s = "www.taobao.com"
		reverseUrl(s)
		fmt.Println(s)
	})
	t.Run("myAtoi", func(t *testing.T) {
		fmt.Println(myAtoi("124232"))
	})
}
