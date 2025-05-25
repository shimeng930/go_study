package base

import (
	"fmt"
)

// Go语言中字符串可以进行循环，使用下表循环获取的Ascii字符，使用range循环获取的unicode字符
func strLoop(s string) {
	// unicode字符的值
	for k, v := range s {
		fmt.Printf("[idx=%d, val=%v(%c), type=%T] ", k, v, v, v)
	}
	fmt.Println()
	// 下标 显示Ascii码
	for i := 0; i < len(s); i++ {
		fmt.Printf("[idx=%d, val=%v(%c), type=%T] ", i, s[i], s[i], s[i])
	}
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%d--%v(%c) ", i, s[i], s[i])
	}
	fmt.Println()

	for i, r := range s { //rune
		fmt.Printf("%d--%v(%c) ", i, r, r)
	}
}

func Strlen(s string) {
	fmt.Printf("字符串字节长度|str=%s, len=%d \n", s, len(s))

	// 字符串长度使用 utf8.RuneCountInString(s) 或者 len([]rune(s))
	fmt.Printf("字符串字符长度|str=%s, len=%d \n", s, len([]rune(s)))
}
