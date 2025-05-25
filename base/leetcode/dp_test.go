package leetcode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_dp(t *testing.T) {
	t.Run("coinChange", func(t *testing.T) {
		fmt.Println(coinChange([]int{2}, 3))
		fmt.Println(coinChange1([]int{1, 2, 5}, 11))
	})
	t.Run("coinChangeAll", func(t *testing.T) {
		fmt.Println(coinChangesAll([]int{1, 2}, 5))
		fmt.Println(coinChangesAll([]int{1, 1, 1, 1, 1}, 4))
		assert.Equal(t, climbStairsV1([]int{1, 2}, 5), climbStairs(5))
	})
	t.Run("wordBreak", func(t *testing.T) {
		fmt.Println(wordBreak1("dogs", []string{"dog", "s", "gs"}))
		fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat", "san"}))
	})
	t.Run("longestPalindrome", func(t *testing.T) {
		fmt.Println(longestPalindromeV("aaaa"))
		fmt.Println(longestPalindrome("bcbacd"))
		fmt.Println(longestPalindromeV1("babad"))
	})
	t.Run("canPartition", func(t *testing.T) {
		fmt.Println(canPartition([]int{3, 3, 3, 4, 5}))
		fmt.Println(canPartition([]int{2, 2, 1, 1}))
		fmt.Println(canPartition([]int{1, 55, 11, 5}))
		fmt.Println(canPartitionDFS([]int{2, 2, 1, 1}))
	})
	t.Run("minPathSum", func(t *testing.T) {
		// [1,3,1],[1,5,1],[4,2,1]
		grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
		fmt.Println(minPathSum(grid))
	})
	t.Run("longestCommonSubsequence", func(t *testing.T) {
		//fmt.Println(longestCommonSubsequence("xaxx", "a"))
		fmt.Println(longestCommonSubsequence("abcde", "ace"))
	})
	t.Run("numDecodingDP", func(t *testing.T) {
		fmt.Println(numDecodingDP("27"))
	})
	t.Run("rob", func(t *testing.T) {
		fmt.Println(rob([]int{2, 7, 9, 3, 1}))
	})
	t.Run("lengthOfLISDP", func(t *testing.T) {
		fmt.Println(lengthOfLISDP([]int{0, 1, 0, 3, 2, 3}))
	})
	t.Run("lengthOfLISDP", func(t *testing.T) {
		a := [...]int{0, 1, 2, 3}
		x := a[:1]
		y := a[2:]
		x = append(x, y...)
		x = append(x, y...)
		fmt.Println(a, x)
	})
}
