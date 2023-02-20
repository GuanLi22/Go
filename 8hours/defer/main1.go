// @Time : 2023/2/15 22:26
// @Author : GuanLi
// Written by Gland

package main

//知识点1:探究 defer 执行顺序

import "fmt"

func func1() {
	fmt.Println("A")
}

func func2() {
	fmt.Println("B")
}

func func3() {
	fmt.Println("C")
}

func main() {
	defer func1()
	defer func2()
	defer func3()
	//遇见defer 依次把他们压入栈 当函数将要结束时 再依次弹出
	//所以打印顺序为 func3 func2 func1
	//详细见当前目录下的图片
}
