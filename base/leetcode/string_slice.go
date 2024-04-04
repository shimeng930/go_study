package leetcode

import (
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
