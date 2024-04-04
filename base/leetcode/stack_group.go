package leetcode

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

// lc42 接雨水
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
