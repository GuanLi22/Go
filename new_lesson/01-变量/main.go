//@Time : 2023/1/29 21:38
//@Author : GuanLi
//Written by Gland

package main

import (
	"fmt"
)

var studentName = "guanli" //go语言中推荐使用小驼峰命名
var a string = "guanli"
var d float32 = 12.1
var e int = 12
var f bool = true
var b = "guanli" //类型推断
var (
	name = "guanli"
	age  = 15
	time = "2023/1/29"
	x    = 0
	y    = 0
)

/*注:非全局变量 声明后必须使用*/

// 定义一个函数
func clac() (a int, b int) {
	return 1, 2
}

func main() {
	c := 12 // 简短变量声明 类似于局部变量 只能在函数内使用
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(time)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(studentName)
	fmt.Printf("name:%v\n", name)

	x, _ = clac()
	_, y = clac() //_为匿名变量
	println(x)
	println(y)
}
