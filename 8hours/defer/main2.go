//@Time : 2023/2/16 7:32
//@Author : GuanLi
//Written by Gland

package main

import "fmt"

// 知识点2: defer和return谁先谁后

func deferFunc() int {
	fmt.Println("deferFunc called..")
	return 0
}

func returnFunc() int {
	fmt.Println("returnFunc called..")
	return 0
}

func returnAndDefer() int {
	defer deferFunc()
	return returnFunc()
}

func main() {
	returnAndDefer()
}

//结论:return先 defer最后
