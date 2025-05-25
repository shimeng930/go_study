package leetcode

import (
	"sort"
)

// SHeap ------------------------------------------------
// 初始化堆的时候，需要把数组0当成虚拟数据；这是为下标计算的时候方便，从1开始即可
type SHeap struct {
	arr []int
}

func NewSHeap(arr []int) *SHeap {
	return &SHeap{arr: arr}
}

func (s *SHeap) size() int {
	return len(s.arr) - 1
}

func (s *SHeap) BuildHeap() *SHeap {
	for i := s.size() / 2; i >= 1; i-- {
		s.shiftDown(i)
	}
	return s
}

func (s *SHeap) swap(i, j int) {
	s.arr[i], s.arr[j] = s.arr[j], s.arr[i]
}

// parent < kid
func (s *SHeap) less(i, j int) bool {
	return s.arr[i] < s.arr[j]
}

// parent > kid
func (s *SHeap) large(i, j int) bool {
	return s.arr[i] > s.arr[j]
}

func (s *SHeap) shiftDown(i int) {
	var flag = true
	var t = i
	for i*2 <= s.size() && flag {
		if s.large(i, i*2) {
			t = i * 2
		}
		if i*2+1 <= s.size() {
			if s.large(t, i*2+1) {
				t = i*2 + 1
			}
		}
		if t == i {
			flag = false
		} else {
			s.swap(i, t)
			i = t
		}
	}

}

func (s *SHeap) push(v int) {
	if v <= s.arr[1] {
		return
	}

	s.arr[1] = v
	s.shiftDown(1)
}

// lc215 数组中第k大的数
func findKthLargestV(nums []int, k int) int {
	karr := []int{0} // 为了下标方便计算，空一个节点出来，实际的根节点从1开始
	karr = append(karr, nums[:k]...)
	sh := NewSHeap(karr).BuildHeap()
	for i := k; i < len(nums); i++ {
		sh.push(nums[i])
	}
	return sh.arr[1]
}

// lc215 数组中第k大的数
func findKthLargest(nums []int, k int) int {
	h := initHeap(nums[0 : k+1])
	for i := k; i < len(nums); i++ {
		if nums[i] < h[1] {
			continue
		}
		h[1] = nums[i]
		shift(h, 1)
	}
	return h[1]
}

func initHeap(nums []int) []int {
	// 0, 1 2,3  4,5 6，7
	var heap = []int{0}
	heap = append(heap, nums...)
	for i := len(nums) / 2; i > 0; i-- {
		shift(heap, i)
	}
	return heap
}

func shift(nums []int, i int) {
	var size = len(nums)
	var nidx = i
	var next = true
	for next {
		if i*2 < size && nums[nidx] > nums[i*2] {
			nidx = i * 2
		}
		if i*2+1 < size && nums[nidx] > nums[i*2+1] {
			nidx = i*2 + 1
		}
		if i == nidx {
			next = false
		} else {
			nums[i], nums[nidx] = nums[nidx], nums[i]
			i = nidx
		}
	}
}

// 给定数组，每次操作一个元素，把最大值变成第二大元素值，重复操作，直到最后元素都相等，返回次数
// 输入：[4，5，5，2，4]：
// 5变成4，需要2次操作；
// 4变成2，需要4次操作 全变成2之后，返回操作次数2+4=6
func maxOperations(nums []int) int {
	var uniqueNums []int
	var numCnt = make(map[int]int)
	for _, n := range nums {
		if _, ok := numCnt[n]; !ok {
			uniqueNums = append(uniqueNums, n)
		}
		numCnt[n]++
	}
	sort.Slice(uniqueNums, func(i, j int) bool {
		return uniqueNums[i] > uniqueNums[j]
	})

	var res, pre int
	for _, n := range uniqueNums {
		res += pre
		pre += numCnt[n]
	}
	return res
}

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

// lc300 最长递增子序列
func lengthOfLIS(nums []int) int {
	var res int
	var f = make([]int, len(nums))
	for i, n := range nums {
		f[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < n {
				f[i] = max(f[i], f[j]+1)
			}
		}
		res = max(res, f[i])
	}
	return res
}

// lc438 找到字符串中所有字母异位词
func findAnagrams(s string, p string) []int {
	var ps = len(p)
	//var pm = make(map[rune]int)
	//for _, c := range p {
	//	pm[c]++
	//}

	//var isAna func(i, j int) bool
	//isAna = func(i, j int) bool {
	//	var ijm = make(map[rune]int)
	//	for x := i; x <= j; x++ {
	//		ijm[rune(s[x])]++
	//	}
	//	for c, n := range pm {
	//		if ijm[c] != n {
	//			return false
	//		}
	//	}
	//	return true
	//}

	var np = [26]int{}
	for _, c := range p {
		np[c-'a']++
	}

	var res []int
	var ns = [26]int{}
	for l, r := 0, 0; r < len(s); r++ {
		ns[rune(s[r])-'a']++
		if r-l < ps-1 {
			continue
		}
		//if isAna(l, r) {
		//	res = append(res, l)
		//}
		if np == ns { // 比较两个数组的效率 比每次比较窗口里的map效率更高
			res = append(res, l)
		}
		ns[rune(s[l])-'a']--
		l++
	}
	return res
}
