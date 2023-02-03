//@Time : 2023/2/3 18:18
//@Author : GuanLi
//Written by Gland

package main

import "fmt"

var (
	a    = 1
	b    = 1.0
	c    = false
	d    = "ha"
	data = "Hello沙河小王子"
)

func main() {
	fmt.Printf("%T %T %T %T ", a, b, c, d)
}
