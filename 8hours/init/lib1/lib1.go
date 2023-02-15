// @Time : 2023/2/15 15:25
// @Author : GuanLi
// Written by Gland
package lib1

import "fmt"

// Lib1包提供的API
func Lib1Test() {
	//函数名大写表示函数可以被外部调用 小写只能在当前包内使用
	fmt.Println("lib1. Lib1Test() ....")
}

func init() {
	fmt.Println("lib1. init() ....")
}
