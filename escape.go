package main

import "fmt"

// 逃逸分析的好处：
// 减少gc的压力，不逃逸的对象分配在栈上，当函数返回时就回收了资源，不需要gc标记清除
// 逃逸分析完后可以确定哪些变量可以分配在栈上，栈的分配比堆快，性能好
// 同步消除，如果你定义的对象的方法上有同步锁，但在运行时，却只有一个线程在访问，此时逃逸分析后的机器码，会去掉同步锁运行

type PointS struct {
	M *int
}

type S struct {
	M int
}

func leakParamsRef (x *int) (z *PointS) {
	//z.M = x
	fmt.Println(x)
	return z
}

//func paramsRef (y int) (z PointS) {
//	z.M = &y
//	y = 2
//	return z
//}
//
func leakParamsRefNoPoint (x *int) (z *S) {
	z.M = *x
	fmt.Println(x)
	return z
}
//
//func ref(y *int, z *PointS) *PointS {
//	z.M = y
//	fmt.Println(z)
//	return z
//}

func Amain() {
	var i int
	//s := S{M:1}

	//var ps PointS
	//fmt.Println(s)
	//leakParamsRef(&i)
	//fmt.Println(s)
	//paramsRef(i)
	//fmt.Println(i)
	//
	//ref(&i, &ps)
	//
	leakParamsRefNoPoint(&i)
}