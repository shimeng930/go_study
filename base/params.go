package base

import "fmt"

func ParamsTest()  {
	var i = 10
	a := &i
	fmt.Printf("i point is %p\n", &a)
	print(&i)
}

func print(i *int)  {
	fmt.Printf("param point is %p\n", &i)
}