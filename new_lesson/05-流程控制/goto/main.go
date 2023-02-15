//@Time : 2023/2/4 13:14
//@Author : GuanLi
//Written by Gland

package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'B' {
				break
				//设置退出标签
				//goto breakTag
			}
			fmt.Printf("%v-%c\n", i, j)
		}
	}
	//return
	// 标签
	//breakTag:
	//	fmt.Println("结束for循环")
}
