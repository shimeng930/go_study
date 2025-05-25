package leetcode

import (
	"math"
	"sort"
)

// lc-33 搜索旋转排序数组
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

// lc-704 二分查找
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

// lc-35
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

func findMedianSortedArrays(nums1, nums2 []int) float64 {
	var nums []int
	var i, j int
	for i < len(nums1) || j < len(nums2) {
		if i == len(nums1) {
			nums = append(nums, nums2[j:]...)
			break
		}
		if j == len(nums2) {
			nums = append(nums, nums1[i:]...)
			break
		}

		if nums1[i] < nums2[j] {
			nums = append(nums, nums1[i])
			i++
		} else {
			nums = append(nums, nums2[j])
			j++
		}
	}

	var l, r = 0, len(nums) - 1
	for l <= r {
		if r-l <= 1 {
			break
		}
		l++
		r--
	}
	return float64(nums[l]+nums[r]) / 2
}

func findMedianSortedArraysV2(nums1, nums2 []int) float64 {
	var fun = func(n1, n2 []int) float64 {
		var m, n = len(n1), len(n2)
		var l, r, leftItemLen = 0, m, (m + n + 1) / 2 // tips: 有边界是m而不是m-1
		for l <= r {
			mid := (l + r) / 2
			n2mid := leftItemLen - mid
			// 找到两个数组分割点左右的值
			// left1 <= right2 && left2 <= right1
			var left1, left2, right1, right2 int
			if mid-1 < 0 {
				left1 = int(math.MinInt32)
			} else {
				left1 = n1[mid-1]
			}
			if mid < m {
				right1 = n1[mid]
			} else {
				right1 = int(math.MaxInt32)
			}

			if n2mid == 0 {
				left2 = math.MinInt32
			} else {
				left2 = n2[n2mid-1]
			}
			if n2mid == n {
				right2 = math.MaxInt32
			} else {
				right2 = n2[n2mid]
			}

			if left1 <= right2 && left2 <= right1 {
				if (m+n)%2 == 1 {
					return float64(max(left1, left2))
				} else {
					return float64(max(left1, left2)+min(right1, right2)) / 2
				}
			} else if left1 > right2 {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		return float64(0)
	}
	if len(nums1) > len(nums2) {
		return fun(nums2, nums1)
	} else {
		return fun(nums1, nums2)
	}
}

// lc.239 滑动窗口最大值
func maxSlidingWindow(nums []int, k int) []int {
	var l, r int
	var res []int
	for r < len(nums) {
		for r-l < k-1 {
			r++
			//if r == len(nums) {
			//	return res
			//}
			for i := 1; i <= r-l && nums[r] > nums[r-i]; i++ {
				nums[r-i] = nums[r]
			}
		}
		if r-l == k-1 {
			res = append(res, nums[l])
		}
		if r == len(nums)-1 {
			break
		}
		l++
	}
	return res
}

// max value in window
func maxSlidingWindowV1(nums []int, k int) []int {
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

// lc15 三数之和
func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		t := 0 - nums[i]
		for j, k := i+1, len(nums)-1; j < k; {
			if nums[j]+nums[k] == t {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				i++
				j--
			} else if nums[j]+nums[k] < t {
				j++
			} else {
				k--
			}
		}
	}
	return res
}
