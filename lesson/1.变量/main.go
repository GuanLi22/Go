//@Time : 2023/1/19 18:51
//@Author : GuanLi
//Written by Gland

package main

import "fmt"

func main() {
	//*变量可以单个声明
	/* var name string
	var age int
	var ma bool
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	fmt.Printf("ma: %v\n", ma) */

	//*可以批量声明
	/* var (
	   	name string
	   	age int
	   	ma bool
	)
	name = "guanli"
	age = 16
	ma = false
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	fmt.Printf("ma: %v\n", ma) */

	//*这样可以设置变量的初始值
	/* var name string = "guanli"
	var age int = 16
	var ma bool = false
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	fmt.Printf("ma: %v\n", ma) */

	//*也可以不声明变量的类型 自动识别类型 类型推导
	/* var name = "guanli"
	var age = 16
	var ma = false
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	fmt.Printf("ma: %v\n", ma) */

	//*一次初始化多个变量
	/* var name, age, ma = "guanli", 16, false
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	fmt.Printf("ma: %v\n", ma) */

	//*短变量声明,指的是在*函数内部*,可以使用:=进行声明和初始化
	/* name := "guanli"
	age := 16
	ma := false
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	fmt.Printf("ma: %v\n", ma) */
	//!注 只能在函数内部使用//
}

// *匿名变量
func getNameAndAge() (string, int) {
	return "guanli", 16
}
func main1() { //应为main
	name, age := getNameAndAge()
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
}
