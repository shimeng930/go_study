package leetcode

import (
	"fmt"
	"strings"
)

// lc 322 coin change
func coinChange(coins []int, amount int) int {

	//var s = "hello"
	//for i:=0;i<len(s);i++ {
	//	v := s[i:]
	//}

	var dfs func(amt int) int
	dfs = func(amt int) int {
		if amt == 0 {
			return 0
		}
		if amt < 0 {
			return -1
		}
		var res = amt + 1
		for _, c := range coins {
			x := dfs(amt - c)
			if x == -1 {
				continue
			}
			//if dfs(amt-c) == -1 {
			//	continue
			//}
			res = min(res, x+1)
		}
		if res == amt+1 {
			return -1
		} else {
			return res
		}
	}

	return dfs(amount)

}

func dfs(amt int, coins []int) int {
	if amt == 0 {
		return 0
	}
	if amt < 0 {
		return -1
	}
	var res = amt + 1
	for _, c := range coins {
		x := dfs(amt-c, coins)
		if x == -1 {
			continue
		}
		//if dfs(amt-c) == -1 {
		//	continue
		//}
		res = min(res, x+1)
	}
	if res == amt+1 {
		return -1
	} else {
		return res
	}
}

func coinChange1(coins []int, amount int) int {
	var dp = make([]int, amount+1)
	// dp[0]=0, 因为金额为0时，只需要0个硬币
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
		for _, c := range coins {
			if i-c < 0 {
				continue
			}
			dp[i] = min(dp[i], dp[i-c]+1)
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

// lc 139 word break
func wordBreak(s string, wordDict []string) bool {

	wordDictSet := make(map[string]bool)
	for _, w := range wordDict {
		wordDictSet[w] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			w := s[j:i]
			if dp[j] && wordDictSet[w] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

func wordBreak1(s string, wordDict []string) bool {
	var dp = make([]bool, len(s)+1)
	dp[0] = true
	for i := 0; i < len(s); i++ {
		for _, w := range wordDict {
			if i+1 < len(w) {
				continue
			}
			if s[i+1-len(w):i+1] == w && !dp[i+1] {
				dp[i+1] = dp[i-len(w)+1]
			}
		}
	}
	return dp[len(s)]
}

// lc5 最长回文
func longestPalindrome(s string) string {
	var size = len(s)
	var dp = make([][]bool, size)
	for i := range dp {
		dp[i] = make([]bool, size)
	}
	var l, r, curMax int
	for i := size - 1; i >= 0; i-- {
		for j := i; j < size; j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1] && s[i+1] == s[j-1]
				}
			}
			if dp[i][j] && j-i+1 > curMax {
				curMax = j - i + 1
				l = i
				r = j
			}
		}
	}
	return s[l : r+1]
}
func longestPalindromeV1(s string) string {
	var tem, l, r int
	// var cl, cr int
	var res string
	for i := 0; i < len(s); i++ {
		curMax := 1
		l, r = i-1, i+1
		for ; l >= 0 && s[l] == s[i]; l-- {
			curMax++
		}
		for ; r < len(s) && s[r] == s[i]; r++ {
			curMax++
		}
		for l >= 0 && r < len(s) && s[l] == s[r] {
			curMax += 2
			l--
			r++
		}
		if curMax > tem {
			tem = curMax
			res = s[l+1 : r]
		}
	}
	return res
}

// lc 416 分割等和子集
func canPartition(nums []int) bool {
	var sum int
	for _, n := range nums {
		sum += n
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	var dp = make([]bool, target+1)
	dp[0] = true
	for _, n := range nums {
		for j := target; j >= n; j-- {
			dp[j] = dp[j] || dp[j-n]
		}
	}
	return dp[target]
}

func canPartitionDFS(nums []int) bool {
	var lsum, rsum int
	var vi = make(map[int]int)
	for _, n := range nums {
		vi[n]++
	}
	var vicnt int
	var res bool

	var dfs func()
	dfs = func() {
		if res {
			return
		}
		if vicnt == len(nums) {
			res = lsum == rsum
			return
		}

		for _, n := range nums {
			if vi[n] == 0 {
				continue
			}

			if lsum > rsum {
				rsum += n
			} else {
				lsum += n
			}
			vi[n]--
			vicnt++
			dfs()
			vi[n]++
			vicnt--
		}
	}
	dfs()
	return res
}

// lc 64
func minPathSum(grid [][]int) int {
	// dp[i][j] = max(dp[i-1][j], dp[i][j-1])
	for i := 1; i < len(grid); i++ {
		grid[i][0] = grid[i][0] + grid[i-1][0]
	}
	for j := 1; j < len(grid[0]); j++ {
		grid[0][j] = grid[0][j] + grid[0][j-1]
	}
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[i]); j++ {
			grid[i][j] = min(grid[i-1][j], grid[i][j-1]) + grid[i][j]
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

func longestCommonSubsequence(text1 string, text2 string) int {
	// s[i,j] = max(s[i, j-1], s[i-1,j]) s[i]!=t[j]
	// s[i,j]=s[i-1,j-1]+1, s[i]==t[j]
	var l1, l2 = len(text1), len(text2)
	if l1 == 0 || l2 == 0 {
		return 0
	}
	var dp = make([][]int, l1)
	for i := 0; i < l2; i++ {
		if strings.ContainsRune(text2, rune(text1[0])) {
			dp[0] = append(dp[0], 1)
		} else {
			dp[0] = append(dp[0], 0)
		}
	}
	for i := 0; i < l1; i++ {
		dp[i] = make([]int, l2)
		if strings.ContainsRune(text1, rune(text2[0])) {
			dp[i][0] = 1
		}
	}
	for i := 1; i < l1; i++ {
		for j := 1; j < l2; j++ {
			if text1[i] == text2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	fmt.Println(dp) // [[1 0 0] [0 0 0] [0 1 1] [0 1 1] [0 1 2]]
	return dp[l1-1][l2-1]
	//var dp = make([][]int, l1+1)
	//for i:=0;i<=l1;i++ {
	//    dp[i]=make([]int, l2+1)
	//}
	//for i, x := range text1 {
	//    for j, y := range text2 {
	//        if x==y {
	//            dp[i+1][j+1]=dp[i][j]+1
	//        }else {
	//            dp[i+1][j+1]=max(dp[i+1][j],dp[i][j+1])
	//        }
	//    }
	//}
	//fmt.Println(dp) // [[0 0 0 0] [0 1 1 1] [0 1 1 1] [0 1 2 2] [0 1 2 2] [0 1 2 3]]
	//return dp[l1][l2]
}

// lc91 解码方法
func numDecodingDP(s string) int {
	// dp[i]是以i为结尾的字符串可能的编码个数
	// if dp[i-1]>2 || {dp[i-1]dp[i]} > 26 dp[i]=dp[i-1]
	// else dp[i]=dp[i-1]+1

	var dp = make([]int, len(s))
	var res = 1
	for i := range s {
		if i == 0 {
			dp[0] = 1
			continue
		}
		if s[i-1] > '2' || s[i-1:i+1] > "26" {
			dp[i] = dp[i-1]
		} else {
			dp[i] = dp[i-1] + 1
		}
		res = max(res, dp[i])
	}
	return res
}

func rob(nums []int) int {
	// dp[i]: rob the index i room, all amount
	// if v[i-1] >= v[i]: dp[i] = dp[i-1]
	// if v[i-1] < v[i] dp[i] = dp[i-2] + v[i]
	if len(nums) == 0 {
		return 0
	}
	var dp []int
	for _, num := range nums {
		dp = append(dp, num)
	}
	if len(nums) == 1 {
		dp[0] = nums[0]
		return nums[0]
	}

	var res int
	for i := 1; i < len(nums); i++ {
		if i == 1 {
			dp[i] = max(dp[0], dp[i])
			res = max(res, dp[i])
			continue
		}
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
		res = max(res, dp[i])
	}
	return res
}

// lc.300
func lengthOfLISDP(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	var res int
	var dp = make([]int, len(nums))
	for i, n := range nums {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if n > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}
