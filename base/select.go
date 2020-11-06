package base

import (
	"fmt"
	"time"
)

func SelectCaseOne() {
	//t1 := time.Tick(time.Second)
	//t2 := time.Tick(time.Second)


	var count int
	for {
		select {
		case <- time.Tick(time.Millisecond * 500):
			fmt.Printf("case one running... count=%d\n", count)
			if count == 4 {
				goto done
			}
			count++
		case <- time.Tick(time.Millisecond*400):
			fmt.Printf("case two running... count=%d\n", count)
			count++
		}
	}

	done:
		fmt.Println("test finish")
}
