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

// lc-33
func findTarget(nums []int, target int) int {
	var l, r = 0, len(nums) - 1
	for l <= r {
		var mid = (l + r) / 2
		if target == nums[mid] {
			return mid
		}

		// 判断有序
		if nums[mid] > nums[l] {
			// 左边有序
			if target >= nums[l] && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		if nums[mid] < nums[l] {
			// 右边有序
			if target > nums[mid] && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}

// lc-27
func removeElement(nums []int, val int) int {
	var l, r, k = 0, 1, 0
	if nums[0] != val {
		k++
	}
	for r < len(nums) {
		if nums[r] == val {
			if nums[l] != val {
				l = r
			}
		} else {
			k++
			for nums[l] != val && l < r {
				l++
			}
			if l < r {
				nums[l] = nums[r]
				nums[r] = val
				l++
			}
		}
		r++
	}
	return k
}

func removeDuplicates(nums []int) int {
	var f, s, l = 2, 2, len(nums)
	if l < 3 {
		return l
	}
	for f < l {
		if nums[s-2] != nums[f] {
			nums[s] = nums[f]
			s++

		}
		f++
	}
	return s
}

func findMax(nums []int) int {
	var l, r = 0, len(nums) - 1
	for l < r {
		var mid = l + (r-l+1)/2
		if nums[mid] < nums[l] {
			r = mid - 1 // 求最大值，因此mid肯定不是最大值，因此r=mid-1
		} else {
			l = mid
		}
	}
	return nums[l]
	//return nums[(r + 1) % len(nums)];    // 这一行是求最小值：最大值向右移动一位就是最小值了（需要考虑最大值在最右边的情况)
}

func findMin(nums []int) int {
	var l, r = 0, len(nums) - 1
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] > nums[r] {
			l = mid + 1 // 求最小值，因此mid肯定不是最小值，因此l=mid+1
		} else {
			r = mid
		}
	}
	return nums[l]
}
