package leetcode

import "fmt"

// leetCode 135.分发糖果
func candy(ratings []int) int {
	var c = make([]int, len(ratings))
	for i, _ := range ratings {
		if i == 0 {
			c[i] = 1
			continue
		}
		if ratings[i] > ratings[i-1] {
			c[i] = c[i-1] + 1
		} else {
			c[i] = 1
		}
	}
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			c[i] = max(c[i], c[i+1]+1)
		}
	}
	var res int
	for _, cnt := range c {
		res += cnt
	}
	return res
}

func jump(nums []int) int {
	// dp[i] 表示到达索引 i 需要的最小跳跃次数
	// var dp = make([]int, len(nums))
	// dp[0]=0
	// for i:=1; i<len(nums); i++ {
	//     dp[i]=i
	//     for j:=0;j<i;j++{
	//         if j+nums[j] >= i {
	//             dp[i] = min(dp[j]+1, dp[i])
	//         }
	//     }
	// }
	// return dp[len(nums)-1]

	// 贪心算法 贪心算法的思路是，在每一步中尽可能跳得更远。
	var maxReach, step, end int
	for i := 0; i < len(nums)-1; i++ {
		fmt.Println(i)
		maxReach = max(nums[i]+i, maxReach)
		if i == end {
			step++
			end = maxReach
		}
		if maxReach >= len(nums)-1 {
			if end != maxReach {
				step++
			}
			break
		}
	}
	return step
}
