package leetcode

import (
	"fmt"
	"sort"
)

// 最长连续序列
// sort and traverse
func longestConsecutive(nums []int) int {
	sort.Ints(nums)
	var res, curLen = 0, 1
	for i, j := 0, 1; j < len(nums) && i < j; {
		var gap = nums[j] - nums[i]
		if gap == 1 {
			curLen++
		} else if gap > 1 || gap < 0 {
			res = max(res, curLen)
			curLen = 0
		}
		i = j
		j++
	}
	res = max(res, curLen)
	return res
}

// use hash
func longestConsecutiveV2(nums []int) int {
	var numMap = make(map[int]struct{})
	for _, item := range nums {
		numMap[item] = struct{}{}
	}

	var res int
	for num, _ := range numMap {
		var cur int
		if _, ok := numMap[num-1]; ok {
			continue
		} else {
			cur = 1
		}
		for {
			if _, ok := numMap[num+1]; ok {
				cur++
				num++
			} else {
				break
			}
		}
		res = max(res, cur)
	}
	return res
}

// ------------------------------------------------
// sum=k 的子数组
func subarraySum(nums []int, k int) int {
	var sum, cnt int
	var preSum = make(map[int]int)
	preSum[0] = 1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if _, ok := preSum[sum-k]; ok {
			cnt += preSum[sum-k]
		}
		preSum[sum]++
	}
	return cnt
}

// ------------------------------------------------
// max value in window
func maxSlidingWindow(nums []int, k int) []int {
	var res []int
	var q = make([]int, len(nums)+1)
	q[0] = 0
	var h, t = 1, 0
	for i := 0; i < len(nums); i++ {
		// replace; here is new item is larger
		for h <= t && nums[i] > nums[q[t]] {
			t--
		}
		// push new item
		t++
		q[t] = i

		// head pop current result
		if q[h] <= i-k {
			h++
		}
		if i+1 >= k {
			res = append(res, nums[q[h]])
		}
	}
	return res
}

// binary search
func binarySearch(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 1
		} else {
			return -1
		}
	}

	if nums[len(nums)/2] > target {
		return binarySearch(nums[:len(nums)/2], target)
	} else if nums[len(nums)/2] == target {
		return 1
	} else {
		return binarySearch(nums[len(nums)/2+1:], target)
	}
}

func searchInsert(nums []int, target int) int {
	//  if len(nums) == 1 {
	// 	return 0
	// }
	var binarySearch func(l, r int) int
	binarySearch = func(l, r int) int {
		if r < l {
			return l
		}
		mid := (r + l) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			if r == l {
				return r
			} else {
				return binarySearch(l, mid-1)
			}
		} else {
			if r == l {
				return r + 1
			} else {
				return binarySearch(mid+1, r)
			}
		}
	}

	return max(0, binarySearch(0, len(nums)-1))
}

func partitionLabels(s string) []int {
	var charMap = make(map[rune][2]int)
	for i, c := range s {
		if _, ok := charMap[c]; ok {
			ci := charMap[c]
			ci[1] = i
			charMap[c] = ci
		} else {
			charMap[c] = [2]int{i, i}
		}
	}
	var lmin, lmax int
	var res []int
	for i, c := range s {
		ci := charMap[c]
		lmin = min(lmin, ci[0])
		lmax = max(lmax, ci[1])
		if i == lmax {
			res = append(res, i+1)
			// lmax=lmin
		}
	}
	for i := len(res) - 1; i > 0; i-- {
		res[i] = res[i] - res[i-1]
	}
	// eaaaabaaec
	return res
}

// lc 91. 解码方法
func numDecodings(s string) int {
	var path []rune
	var dfs func(data string)
	var cnt int
	dfs = func(data string) {
		if len(path) > 0 && (len(path) < 2 || string(path) < "26") {
			cnt++
		}
		if len(data) == 0 {
			return
		}
		for i, c := range data {
			path = append(path, c)
			dfs(data[i+1:])
			path = path[:len(path)-1]
		}
	}
	dfs(s)
	return cnt
}

// 76. 最小覆盖子串 滑动窗口
func minWindow(s string, t string) string {
	var wind = make(map[byte]int)
	var need = make(map[byte]int)
	for _, c := range t {
		need[byte(c)]++
	}

	var l, r = 0, 0
	var res string
	for r < len(s) {
		// add into wind
		wind[s[r]]++
		fmt.Println(s[l : r+1])
		for containsAll(wind, need) {
			if len(res) > (r-l) || len(res) == 0 {
				res = s[l : r+1]
			}

			wind[s[l]]--
			l++
		}
		r++
	}
	return res
}

func containsAll(a, b map[byte]int) bool {
	for k, v := range b {
		if a[k] < v {
			return false
		}
	}
	return true
}

// 567. 字符串的排列
func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}

	var wind, need = make(map[byte]int), make(map[byte]int)
	for _, c := range s1 {
		need[byte(c)]++
	}

	var l, r, ws = 0, 0, 0
	for r < len(s2) {
		for ws < len(s1) {
			wind[s2[r]]++
			ws++
			r++
		}

		if containsAll(wind, need) {
			return true
		}

		wind[s2[l]]--
		l++
		ws--
	}
	return false
}

func lengthOfLongestSubstring(s string) int {
	var l, r, res = 0, 0, 0
	var wind = make(map[byte]int)

	for r < len(s) {
		if old, ok := wind[s[r]]; ok {
			l = max(l, old+1)
		}

		wind[s[r]] = r
		r++

		res = max(res, r-l)

	}
	return res
}

// lc33 搜索旋转排序数组
func searchArr(nums []int, target int) int {
	var l, r = 0, len(nums) - 1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[l] <= nums[mid] {
			if nums[l] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}

// lc8. 字符串转换整数 (atoi)
func myAtoi(s string) int32 {
	var res int32
	var mark bool
	//const maxInt32 = math.MaxInt32
	//const minInt32 = math.MinInt32
	const maxInt32 = 2<<30 - 1
	const minInt32 = -maxInt32 - 1
	for _, c := range s {
		if c == ' ' {
			continue
		}
		if c == '-' {
			mark = true
			continue
		}

		if maxInt32/10 < res || maxInt32/10 == res && maxInt32%10 < int32(c-'0') {
			return maxInt32
		}

		res = (res * 10) + (c - '0')
	}
	if mark {
		res = res * -1
	}
	return res
}
