// @Time : 2023/1/30 16:54
// @Author : GuanLi
// Written by Gland
package main

import "fmt"

func main() {
	// 定义一个十进制的数
	var a1 = 10
	fmt.Printf("%d\n", a1) // 十进制数输出用%d
	fmt.Printf("%b\n", a1) // 十进制数以2进制输出
	fmt.Printf("%o\n", a1) // 十进制数以8进制输出
	fmt.Printf("%x\n", a1) // 十进制数以16进制输出
	// 八进制转十进制
	a2 := 077
	fmt.Printf("%d\n", a2)
	// 十六进制转十进制
	a3 := 0x123456
	fmt.Printf("%d\n", a3)
	/*查看变量的类型*/
	fmt.Printf("%T\n", a1)
	fmt.Printf("%T\n", a2)
	fmt.Printf("%T\n", a3)

	//声明int8类型的变量
	var b1 = int8(9) //要指定声明int8类型否则会声明int类型
	fmt.Printf("%T\n", b1)

}
