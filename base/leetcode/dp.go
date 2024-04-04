package leetcode

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
	for i := 0; i <= amount; i++ {
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
