package main

import (
	"fmt"
	"math"
)

func main() {
	var a = 10 //十进制
	//0开头表示八进制
	var b int = 077
	//0x表示16进制
	var c = 0xff1222

	var d complex128 = 1 + 2i
	var f complex64 = 2 + 3i
	fmt.Println(d)
	fmt.Println(f)
	fmt.Printf("%d\n", a) //%d表示输出十进制
	fmt.Printf("%b\n", a) //%b表示输出二进制
	fmt.Printf("%o\n", b) //%o表示输出八进制
	fmt.Printf("%x\n", c) //%x表示输出十六进制
	fmt.Printf("%X\n", c)
	fmt.Printf("%f\n", math.Pi)   //%f表示输出浮点类型,默认小数点后6位
	fmt.Printf("%.2f\n", math.Pi) //指定小数点后两位
}
