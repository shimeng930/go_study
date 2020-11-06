package base

import "fmt"

func ForGoroutine(num int) []int {
	ch := make(chan int)

	for i:=0; i<num; i++ {
		go func(i int) {
			ch <- i
			fmt.Println(i)
		}(i)
	}

	//go func() {
	//	for i:=0; i<num; i++ {
	//		ch <- i
	//	}
	//}()

	var rsp []int
	for d := range ch {
		rsp = append(rsp, d)
		if len(rsp) == num {
			break
		}
	}
	return rsp
}
