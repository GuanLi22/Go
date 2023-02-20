// @Time : 2023/2/15 22:21
// @Author : GuanLi
// Written by Gland
package main

import "fmt"

func main() {
	//写入defer关键字 这是函数结束前 最后运行的语句
	defer fmt.Println("main end1")
	defer fmt.Println("main end2")

	fmt.Println("main::hello go 1")
	fmt.Println("main::hello go 1")
}
