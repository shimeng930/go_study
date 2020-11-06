package base

import "fmt"

func ChannelTest()  {
	ch := make(chan int)
	ch <- 1

	for data := range ch {
		fmt.Println(data)
		break
	}
}

func ChannelTestOne()  {
	ch := make(chan int, 4)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		// 如果不关闭channel,会引发panic
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)

	}
}
