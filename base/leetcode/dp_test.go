package leetcode

import (
	"fmt"
	"testing"
)

func Test_dp(t *testing.T) {
	t.Run("coinChange", func(t *testing.T) {
		fmt.Println(coinChange([]int{2}, 3))
	})
	t.Run("coinChange", func(t *testing.T) {
		fmt.Println(coinChange1([]int{1, 2, 5}, 11))
	})
	t.Run("wordBreak", func(t *testing.T) {
		fmt.Println(wordBreak1("dogs", []string{"dog", "s", "gs"}))
		fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat", "san"}))
	})
	t.Run("longestPalindrome", func(t *testing.T) {
		fmt.Println(longestPalindrome("babad"))
		fmt.Println(longestPalindromeV1("babad"))
	})
	t.Run("canPartition", func(t *testing.T) {
		fmt.Println(canPartition([]int{2, 2, 1, 1}))
		fmt.Println(canPartition([]int{1, 55, 11, 5}))
	})
}

func Test_back_tracking(t *testing.T) {
	t.Run("letterCombinations", func(t *testing.T) {
		fmt.Println(letterCombinations("23"))
	})
}
