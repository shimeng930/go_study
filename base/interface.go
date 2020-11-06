package base

import "fmt"

type People interface {
	ReturnRole() string
}

// 定义一个结构体
type Student struct {
	Role string
}

// 定义结构体的一个方法
func (s Student) ReturnRole() string {
	return s.Role
}

type Teacher struct {
	Role string
}

func (t *Teacher) ReturnRole() string {
	return t.Role
}

func PrintTest() {
	stu := Student{Role:"student"}
	teacher := Teacher{Role:"teacher"}

	var a People
	// 因为Students实现了接口所以直接赋值没问题
	// 如果没实现会报错：cannot use cbs (type Student) as type People in assignment:Student does not implement People (missing ReturnName method)
	a = stu
	fmt.Println(&a)
	fmt.Println(&stu)
	fmt.Println(a.ReturnRole())
	stu.Role = "new role"
	fmt.Println(a.ReturnRole())
	a = &teacher
	fmt.Println(a.ReturnRole())
	teacher.Role = "new tea"
	fmt.Println(a.ReturnRole())

	fmt.Println(a.(People))

}
