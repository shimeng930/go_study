package leetcode

import (
	"fmt"
	"testing"
)

func Test_NewMinStack(t *testing.T) {
	t.Run("SHeap", func(t *testing.T) {
		reverseString([]byte("hello"))

		res := NewSHeap([]int{0, 99, 5, 36, 7, 22, 17, 46, 12, 2, 19, 25, 28, 1, 92}).BuildHeap()
		fmt.Println(res)

		findKthLargest([]int{-1, 2, 0}, 2)

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
}
