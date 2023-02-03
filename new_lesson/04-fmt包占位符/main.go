//@Time : 2023/2/1 7:57
//@Author : GuanLi
//Written by Gland

package main

import "fmt"

//  fmt包占位符

var a = 66

func main() {
	fmt.Printf("%T\n", a) // 查看类型
	fmt.Printf("%v\n", a) // 查看变量的值(万能)
	fmt.Printf("%d\n", a) // 打印十进制
	fmt.Printf("%b\n", a) // 打印二进制
	fmt.Printf("%o\n", a) // 打印八进制
	fmt.Printf("%x\n", a) // 打印十六进制
	var s = "Hello Go"
	fmt.Printf("%s\n", s)  // 打印字符串
	fmt.Printf("%v\n", s)  // 也可以用%v
	fmt.Printf("%#v\n", s) // 会把引号也打印出来
}
