package leetcode

import (
	"fmt"
	"testing"
)

func Test_slice(t *testing.T) {
	t.Run("coinChange", func(t *testing.T) {
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
		fmt.Println(findTarget([]int{8,9,3,4,5,6,7}, 3))
	})

}
