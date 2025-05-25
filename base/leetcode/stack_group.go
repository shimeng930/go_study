package leetcode

import "strconv"

// lc-155 min stack, use another min-stack slice to do
type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{},
	}
}

func (s *MinStack) Push(val int) {
	s.stack = append(s.stack, val)
	var minVal = val
	if len(s.minStack) > 0 {
		minVal = min(val, s.minStack[len(s.minStack)-1])
	}
	s.minStack = append(s.minStack, minVal)
}

func (s *MinStack) Pop() {
	s.stack = s.stack[:len(s.stack)-1]
	s.minStack = s.minStack[:len(s.minStack)-1]
}

func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1]
}

func (s *MinStack) GetMin() int {
	return s.minStack[len(s.minStack)-1]
}

// another solution: use link-list
type listNode struct {
	val  int
	min  int
	next *listNode
}

type MinStackV1 struct {
	head *listNode
}

func NewMinStackV1() MinStackV1 {
	return MinStackV1{}
}

func (s *MinStackV1) Push(val int) {
	if s.head == nil {
		s.head = &listNode{
			val: val,
			min: val,
		}
		return
	}

	newNode := &listNode{
		val:  val,
		min:  min(val, s.head.min),
		next: s.head,
	}

	s.head = newNode
}

func (s *MinStackV1) Pop() {
	s.head = s.head.next
}

func (s *MinStackV1) Top() int {
	return s.head.val
}

func (s *MinStackV1) GetMin() int {
	return s.head.min
}

// 单调栈
func dailyTemperatures(temperatures []int) []int {
	//
	var size = len(temperatures)
	var res = make([]int, size)
	var st = NewStack(size)
	st.push(0)
	for i, n := range temperatures {
		if i == 0 {
			continue
		}
		for idx := st.getTop(); idx >= 0 && n > temperatures[idx]; {
			res[idx] = i - idx
			st.pop()
			idx = st.getTop()
		}
		st.push(i)
	}

	return res
}

func dailyTemperatureV1(temperatures []int) []int {
	var stack []int
	var res = make([]int, len(temperatures))

	for i, n := range temperatures {
		size := len(stack)
		if size == 0 || n <= temperatures[stack[size-1]] {
			stack = append(stack, i)
			continue
		}
		for size >= 0 {
			if size == 0 || n <= temperatures[stack[size-1]] {
				stack = append(stack, i)
				break
			}

			res[stack[size-1]] = i - stack[size-1] + 1
			stack = stack[:size-1]
			size--
		}

	}
	return res
}

type stack struct {
	arr  []int
	size int
	t    int
}

func NewStack(cap int) *stack {
	return &stack{arr: make([]int, cap), size: cap}
}

func (s *stack) pop() {
	if s.t > 0 {
		s.t--
	}
}

func (s *stack) push(n int) {
	if s.t < s.size {
		s.t++
	}
	s.arr[s.t-1] = n
}

func (s *stack) getTop() int {
	if s.t > 0 {
		return s.arr[s.t-1]
	}
	return -1
}

func dailyTemperatures1(temperatures []int) []int {
	var stack []int
	// mo up; if < top; in; else record res
	var res = make([]int, len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		if len(stack) == 0 {
			stack = append(stack, i)
			continue
		}
		for len(stack) > 0 {
			topIdx := stack[len(stack)-1]
			if temperatures[i] > temperatures[topIdx] {
				res[topIdx] = i - topIdx
				// stack[0]=i
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}
		stack = append(stack, i)
	}
	return res
}

// 73, 74, 75, 71, 69, 72, 76, 73
func monotonicStack(arr []int, isAsc bool) {
	var stack = make([]int, 0)
	for i, val := range arr {
		for len(stack) > 0 {
			if isAsc {
				if val > arr[stack[len(stack)-1]] {
					break
				}
			} else {
				if val < arr[stack[len(stack)-1]] {
					break
				}
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
}

// lc42 接雨水 单调栈
func trap(height []int) int {
	// 前后缀
	var pre = make([]int, len(height))
	pre[0] = height[0]
	var suf = make([]int, len(height))
	suf[len(height)-1] = height[len(height)-1]

	for i, j := 1, len(height)-2; j >= 0; i, j = i+1, j-1 {
		pre[i] = max(pre[i-1], height[i])
		suf[j] = max(suf[j+1], height[j])
	}
	// 0,1,0,2,1,0,1,3,2,1,2,1
	var res int
	for i, item := range height {
		var area int
		if pre[i] < suf[i] {
			area = pre[i] - item
		} else {
			area = suf[i] - item
		}
		res += area
		//if area > 0 {
		//}
	}
	return res
}

// lc42 接雨水 双指针
func trapV1(height []int) int {
	var left, right = 0, len(height) - 1
	var leftMax, rightMax = height[left], height[right]
	var res int
	for left < right {
		if leftMax <= rightMax {
			if height[left] < leftMax {
				res += leftMax - height[left]
			}
			left++
			leftMax = max(height[left], leftMax)
		} else {
			if height[right] < rightMax {
				res += rightMax - height[right]
			}
			right--
			rightMax = max(height[right], rightMax)
		}
	}
	return res
}

func decodeString(s string) string {
	var stack []string
	var strtmp, ntmp []rune
	var res string
	for i, c := range s {
		if c >= '0' && c <= '9' {
			ntmp = append(ntmp, c)
		}
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
			strtmp = append(strtmp, c)
		}

		if c == '[' {
			stack = append(stack, string(strtmp))
			strtmp = []rune{}
			stack = append(stack, string(ntmp))
			ntmp = []rune{}
		}
		if c == ']' {
			n, _ := strconv.Atoi(stack[len(stack)-1])
			var cur string
			for i := 0; i < n; i++ {
				cur += string(strtmp)
			}
			strtmp = []rune(stack[len(stack)-2] + cur)
			stack = stack[:len(stack)-2]
		}
		if i == len(s)-1 && len(strtmp) > 0 {
			res += string(strtmp)
		}
	}
	return res
}

func decodeString1(s string) string {
	var stack []string
	var strBuilder []rune
	var numBuilder []rune

	for _, c := range s {
		if c >= '0' && c <= '9' {
			numBuilder = append(numBuilder, c)
		} else if c == '[' {
			// 将当前字符串和数字压入栈
			stack = append(stack, string(strBuilder))
			stack = append(stack, string(numBuilder))
			// 重置构建器
			strBuilder = []rune{}
			numBuilder = []rune{}
		} else if c == ']' {
			// 弹出数字和字符串
			numStr := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			prevStr := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			num, _ := strconv.Atoi(numStr)
			repeated := string(strBuilder)
			for i := 0; i < num; i++ {
				prevStr += repeated
			}
			strBuilder = []rune(prevStr)
		} else {
			strBuilder = append(strBuilder, c)
		}
	}

	return string(strBuilder)
}

func decodeWays(s string) {
	var size = len(s)
	var dp = make([]int, size+1)
	for i, c := range s {
		if i == 0 {
			if c == '0' {
				dp[i] = 0
			} else {
				dp[i] = 1
			}
			continue
		}
		if c != '0' {
			dp[i] += dp[i-1]
		}
		if s[i-1] == '1' || s[i-1] == '2' {
			val := int(s[i-1]-'0')*10 + int(c-'0')
			if val <= 26 {
				if i == 1 {
					dp[i]++
				} else {
					dp[i] += dp[i-2]
				}
			}
		}
	}
}

// lc.946 验证栈序列
func validateStackSequences(pushed []int, popped []int) bool {
	var stack = make([]int, len(pushed))
	var topIdx, popIdx = -1, 0
	for _, n := range pushed {
		topIdx++
		stack[topIdx] = n
		for j := popIdx; j < len(popped); j++ {
			if topIdx > -1 && popped[j] == stack[topIdx] {
				topIdx--
				popIdx++
			} else {
				break
			}
		}
	}
	return topIdx == -1
}

func largestRectangleArea(heights []int) int {
	var res, hs = 0, len(heights)
	var rightMin, leftMin = make([]int, hs), make([]int, hs)
	for i := 0; i < hs; i++ {
		leftMin[i] = -1
		rightMin[i] = hs
	}
	var stack []int

	for i, n := range heights {
		size := len(stack)
		if size == 0 || n >= heights[stack[size-1]] {
			stack = append(stack, i)
			continue
		}
		for size > 0 && n < heights[stack[size-1]] {
			rightMin[stack[size-1]] = i
			stack = stack[:size-1]
			size--
		}
		stack = append(stack, i)
	}
	stack = []int{}
	for i := hs - 1; i >= 0; i-- {
		size := len(stack)
		n := heights[i]
		if size == 0 || n >= heights[stack[size-1]] {
			stack = append(stack, i)
			continue
		}
		for size > 0 && n < heights[stack[size-1]] {
			leftMin[stack[size-1]] = i
			stack = stack[:size-1]
			size--
		}
		stack = append(stack, i)
	}

	for i, h := range heights {
		res = max(res, (rightMin[i]-leftMin[i]-1)*h)
	}
	return res
}

// lc1209 删除字符串中所有相邻重复项II
func removeDuplicatesV(s string, k int) string {
	type item struct {
		elem rune
		cnt  int
	}
	var st []*item
	for _, n := range s {
		if len(st) == 0 {
			st = append(st, &item{elem: n, cnt: 1})
			continue
		}

		top := st[len(st)-1]
		if n == top.elem {
			top.cnt++
			if top.cnt == k {
				st = st[:len(st)-1]
			}
		} else {
			st = append(st, &item{elem: n, cnt: 1})
		}
	}
	var nst []rune
	for _, n := range st {
		for i := 0; i < n.cnt; i++ {
			nst = append(nst, n.elem)
		}
	}
	return string(nst)
}
