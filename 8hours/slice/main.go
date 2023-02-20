//@Time : 2023/2/16 11:31
//@Author : GuanLi
//Written by Gland

package main

//切片的介绍

import "fmt"

// 其实是引用传递 也就是传递了指针; 因为是动态的所以数据肯定是在heap上 所以传参会自动传指针
func printArray(myArray []int) {
	for index, value := range myArray {
		fmt.Printf("the index and value of myArray is %v and %v\n", index, value)
	}
	myArray[0] = 999 // 因为是引用传递所以可以直接修改
}

func main() {
	myArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9} //动态数组 也就是切片slice

	fmt.Printf("the type of myArray is %T\n", myArray)
	printArray(myArray)
	fmt.Println(myArray[0])
}
