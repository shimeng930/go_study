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
