package leetcode

// lc 17 phone no
func letterCombinations(digits string) []string {
	var m = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	var res []string
	// dfs(i)->pick i, or not pick. then dfs(i+1)
	var dfs func(idx int)
	var picks []rune
	dfs = func(idx int) {
		if idx == len(digits) {
			var tmp = string(picks)
			res = append(res, tmp)
			return
		}

		// pick
		cidx := digits[idx] - 48
		for _, c := range m[cidx] {
			picks = append(picks, c)
			dfs(idx + 1)
			picks = picks[:len(picks)-1]
		}
	}
	dfs(0)
	return res
}

// lc.46 全排列
func listAll(nums []int) [][]int {
	var res [][]int
	var cur []int
	var dfs func()
	var pick = make(map[int]bool)
	dfs = func() {
		if len(cur) == len(nums) {
			var nc = make([]int, len(nums))
			copy(nc, cur)
			res = append(res, nc)
			return
		}

		for _, n := range nums {
			if pick[n] {
				continue
			}
			// choose
			cur = append(cur, n)
			pick[n] = true
			dfs()
			// revert
			pick[n] = false
			cur = cur[:len(cur)-1]
		}
	}
	dfs()

	return res
}

// 子集和取k个数组成的组合
func listSub(nums []int, k int) [][]int {
	var res [][]int
	var cur []int
	var dfs func(idx int)
	dfs = func(idx int) {
		if k == -1 || len(cur) == k {
			var nc = make([]int, len(cur))
			copy(nc, cur)
			res = append(res, nc)
		}

		for i := idx; i < len(nums); i++ {
			// choose
			cur = append(cur, nums[i])
			dfs(i + 1)
			// revert
			cur = cur[:len(cur)-1]
		}
	}
	dfs(0)

	return res
}

// lc 77 组合
func combine(n int, k int) [][]int {
	var res [][]int
	var path []int
	var dfs func(idx int)
	dfs = func(idx int) {
		if len(path) == k {
			var cp = make([]int, k)
			copy(cp, path)
			res = append(res, cp)
			return
		}

		for i := idx; i <= n; i++ {
			path = append(path, i)
			dfs(i)
			path = path[:len(path)-1]
		}
	}
	dfs(1)
	return res
}

// lc 22 生成括号
func generateParenthesis(n int) []string {
	var res []string
	var path []byte
	var open, closing int
	var bytes = []byte{'(', ')'}
	var dfs func()
	dfs = func() {
		if len(path) == 2*n && open == closing {
			var cp = make([]byte, 2*n)
			copy(cp, path)
			res = append(res, string(cp))
			return
		}

		if open < closing || open > n || closing > n {
			return
		}

		for _, b := range bytes {
			path = append(path, b)
			if b == '(' {
				open++
			} else {
				closing++
			}

			dfs()
			if b == '(' {
				open--
			} else {
				closing--
			}
			path = path[:len(path)-1]
		}
	}
	dfs()
	return res
}
