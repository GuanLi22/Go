// @Time : 2023/2/16 13:23
// @Author : GuanLi
// Written by Gland

package main

import "fmt"

func main() {
	//声明一个slice1切片 变初始化默认值为1,2,3,4,5 长度len为3
	slice1 := []int{1, 2, 3, 4, 5}
	//

	//声明一个slice2切片 但没有给slice2分配空间
	var slice2 []int
	slice2 = make([]int, 3) // 给slice2开辟3个长度 默认值是0
	//make函数 长度指的是切片的长度 容量指的是分配给切片的长度
	//切片:大小指定长度。切片的容量等于它的长度。可以提供第二个整数参数来指定不同的容量;必须不小于长度。
	//例如，make([]int, 0, 10)分配一个大小为10的底层数组，并返回由该底层数组支持的长度为0和容量为10的切片。
	slice2[0] = 100
	//

	//声明一个slice3切片 同时分配3个空间 但是只能初始化值为0
	var slice3 []int = make([]int, 3) // 可以这样和二为一
	//

	//声明一个slice3切片 同时分配3个空间 但是只能初始化值为0 通过类型推断出是切片类型
	slice4 := make([]int, 3)

	//判断slice是否为空
	if slice2 == nil {
		fmt.Println("Empty!")
	}

	fmt.Printf("len = %d, value = %v\n", len(slice1), slice1)
	fmt.Printf("len = %d, value = %v\n", len(slice2), slice2)
	fmt.Printf("len = %d, value = %v\n", len(slice3), slice3)
	fmt.Printf("len = %d, value = %v\n", len(slice4), slice4)
}
