// @Time : 2023/2/3 20:38
// @Author : GuanLi
// Written by Gland
package main

import (
	"fmt"
)

func main() {
	//	for 循环基本格式

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 可以省略初始语句 在外面定义好i

	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}

	// 可以开始结束都省略 但要在循环后写好

	i = 0 // 正常是要写:的
	for i < 10 {
		fmt.Println("i")
		i++
	}

	// 死循环

	/*	for {
		fmt.Println("1")
	}*/

	//for range 循环 遍历

	s := "Hello长春"
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	}

	//跳出循环

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 5 {
			break
		}
	}

	//跳过循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

}
