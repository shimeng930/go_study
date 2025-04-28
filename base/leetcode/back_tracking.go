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

// 全排列
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
