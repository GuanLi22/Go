//@Time : 2023/2/3 20:16
//@Author : GuanLi
//Written by Gland

package main

import "fmt"

//var age = 19

func main() {
	/*	if age > 18 {
			fmt.Println("你已成年")
		} else {
			fmt.Println("你没成年")
		}*/
	// 判断多个条件

	/*	if age >= 35 {
			fmt.Println("中年了")
		} else if age >= 18 {
			fmt.Println("你还是青年")
		} else {
			fmt.Println("你还没成年")
		}*/

	// 特殊写法
	// age 变量只在这个判断里有效
	if age := 19; age > 18 {
		fmt.Println("你已成年")
	} else {
		fmt.Println("你没成年")
	}
}
