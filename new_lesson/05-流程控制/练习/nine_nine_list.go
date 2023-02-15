// @Time : 2023/2/4 10:05
// @Author : GuanLi
// Written by Gland
package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		for v := 1; v <= i; v++ {
			fmt.Printf("%d×%d=%d\t", v, i, v*i)
		}
		fmt.Println("")
	}
	fmt.Println("---------------------------------------------------------------------------")
	for i := 9; i > 0; i-- {
		for v := 1; v <= i; v++ {
			fmt.Printf("%d×%d=%d\t", v, i, v*i)
		}
		fmt.Println("")
	}
}
