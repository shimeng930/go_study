package leetcode

import (
	"fmt"
	"testing"
)

func Test_NewMinStack(t *testing.T) {
	t.Run("SHeap", func(t *testing.T) {
		reverseString([]byte("hello"))

		v1 := NewMinStackV1()
		v1.Push(2147483646)
		v1.Push(2147483646)
		v1.Push(2147483647)
		v1.Pop()
		v1.Pop()
		v1.Push(2147483647)
		v1.Push(-2147483648)
	})
}

func reverseString(s []byte) {
	if len(s) <= 1 {
		return
	}
	s[0], s[len(s)-1] = s[len(s)-1], s[0]
	if len(s) == 2 {
		return
	}
	reverseString(s[1 : len(s)-1])
}

func Test_dailyTemperatures(t *testing.T) {
	t.Run("dailyTemperatures", func(t *testing.T) {
		dailyTemperatures1([]int{73, 74, 75, 71, 69, 72, 76, 73})
		dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})
	})
	t.Run("trap", func(t *testing.T) {
		trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
	})
	t.Run("trapV1", func(t *testing.T) {
		trapV1([]int{4, 2, 0, 3, 2, 5})
	})
}

func Test_dailyTemperaturesV2(t *testing.T) {
	t.Run("monotonicStack", func(t *testing.T) {
		monotonicStack([]int{73, 74, 75, 71, 69, 72, 76, 73}, false)
	})
	t.Run("validateStackSequences", func(t *testing.T) {
		validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1})
	})
	t.Run("decodeString", func(t *testing.T) {
		fmt.Println(decodeString("3[a]2[bc]"))
		fmt.Println(decodeString("3[a2[c]]"))
	})
	t.Run("largestRectangleArea", func(t *testing.T) {
		fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
	})
}
