package leetcode

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
func findKthLargest(nums []int, k int) int {
	karr := []int{0} // 为了下标方便计算，空一个节点出来，实际的根节点从1开始
	karr = append(karr, nums[:k]...)
	sh := NewSHeap(karr).BuildHeap()
	for i := k; i < len(nums); i++ {
		sh.push(nums[i])
	}
	return sh.arr[1]
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
