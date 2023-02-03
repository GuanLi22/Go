// @Time : 2023/2/1 7:37
// @Author : GuanLi
// Written by Gland
package main

import "fmt"

func main() {
	b := true
	var b2 bool //默认是false
	fmt.Printf("%T,%v\n", b2, b2)
	fmt.Printf("%T,%v\n", b, b) // %v万能
}
