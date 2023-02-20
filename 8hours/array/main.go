//@Time : 2023/2/16 8:05
//@Author : GuanLi
//Written by Gland

package main

import "fmt"

// 写[10]int只能传 类型是[10]int的 像[4]int就不能传
func printArray(myArray [10]int) {
	for index, value := range myArray { //值拷贝过程 不能在函数中修改
		fmt.Printf("index = %v, value = %v\n", index, value)
	}
}

func main() {
	//这样定义固定长度的数组
	var myArray [10]int //数组内部的数据类型 和 一共10个

	//可以这样定义数组
	myArray2 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	//遍历数组
	for i, i2 := range myArray2 {
		fmt.Printf("INDEX = %v VALUE = %v\n", i, i2) //这样可以取出索引
	} // i为索引  i2是数组中的值

	//也可以这样遍历
	//for i := 0; i < 10; i++ {
	//}
	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}

	//查看数组的数据类型
	fmt.Printf("myArray1 type = %T\n", myArray)
	fmt.Printf("myArray2 type = %T\n", myArray2)

	//数组传参
	printArray(myArray2)
}
