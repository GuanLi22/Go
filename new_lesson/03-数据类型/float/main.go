// @Time : 2023/1/31 21:21
// @Author : GuanLi
// Written by Gland
package main

import (
	"fmt"
)

// 浮点数
func main() {
	//math.MaxFloat32
	f1 := 1.23456
	fmt.Printf("%T\n", f1) // 默认go语言中的小数都是float64
	var f2 = float32(1.23456)
	fmt.Printf("%T\n", f2) // 声明float32小数
	//f1 = f2 类型不同 不能直接给对方赋值
}
