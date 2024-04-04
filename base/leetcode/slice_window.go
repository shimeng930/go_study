package leetcode

// lc189 轮转数组
func rotate(nums []int, k int) {
	var cnt int
	// []int{-1,-100,3,99}
	for i := 0; i < len(nums); i++ {
		pre, ci := nums[i], i
		ni := (ci + k) % len(nums)
		for ni != i {
			ni = (ci + k) % len(nums)
			tem := nums[ni]
			nums[ni] = pre
			cnt += 1
			ci = ni
			pre = tem
		}
		if cnt == len(nums) {
			break
		}
	}
}

func reverseUrl(url string) {
	var re = func(str string) {
		s := []rune(str)
		for i, n := 0, len(s); i < n/2; i++ {
			s[i], s[n-1-i] = s[n-1-i], s[i]
		}
	}

	var start int
	for i, c := range url {
		if c == '.' {
			re(url[start:i])
			start = i + 1
			continue
		}
	}
}

func lengthOfLIS(nums []int) int {
	var res int
	for i := 0; i < len(nums)-1; {
		j := i + 1
		for ; j < len(nums); j++ {
			if nums[j] <= nums[j-1] {
				res = max(res, j-i)
				i = j
				break
			}
		}
		// res=max(res, j-i+1)
	}
	return res
}
