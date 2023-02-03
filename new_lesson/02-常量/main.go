// @Time : 2023/1/29 23:21
// @Author : GuanLi
// Written by Gland
package main

import "fmt"

// 其定义方法与变量相同
const name string = "guanli" //定义一个常量
const a = "111"
const b = true
const (
	age  = 15
	time = "2023/1/29"
	time1
	time2 //没有赋值和上一行一样
	time3
)

// iota 在 const 关键字出现时将被重置为0。const 中每新增一行常量声明将使 iota 计数一次
// （iota 可理解为 const 语句块中的行索引）。使用 iota 能简化定义，在定义枚举时很有用。
const (
	n1 = iota
	n2
	n3
	n4
)

/*一些iota实例*/
// _ 跳过一些值
const (
	a1 = iota
	a2
	_
	a3
)

// 中间插队
const (
	b1 = iota
	b2 = "guanli"
	b3 = iota
	b4
)

// 一次赋值多个常量
const (
	c1, c2 = iota + 1, iota + 2
	c3, c4 = iota + 1, iota + 2
)

// 定义数量级
// 这里的<<表示左移操作，1<<10表示将1的二进制表示向左移10位，也就是由1变成了10000000000，也就是十进制的1024
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	fmt.Println(a)
	fmt.Println(name)
	fmt.Println(b)
	fmt.Println(age)
	fmt.Println(time)
	fmt.Println(time1)
	fmt.Println(time2)
	fmt.Println(time3)
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b3)
	fmt.Println(b4)
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
	fmt.Println(c4)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)
}
