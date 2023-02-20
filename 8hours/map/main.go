// @Time : 2023/02/16 23:02:52
// @Author : GuanLi
// Written by Visual Studio Code

package main

import "fmt"

func main() {
	//第一种声明方式
	//声明一个map类型数据 key是string类型 value是srting类型
	var myMap map[string]string
	if myMap == nil {
		fmt.Println("myMap 是一个空map")
	}

	//在使用map前 要先给map分配数据空间
	myMap1 := make(map[string]string, 10)

	myMap1["one"] = "java"
	myMap1["two"] = "C++"
	myMap1["three"] = "python"

	fmt.Println(myMap1)

	// 第二种声明方式
	myMap2 := make(map[int]string) // 后面的空间可以不写 会自动分配
	myMap2[1] = "java"
	myMap2[2] = "C++"
	myMap2[3] = "python"

	fmt.Println(myMap2)

	// 第三种方式
	myMap3 := map[string]string{
		"one":   "java",
		"two":   "C++",
		"three": "python",
	}
	fmt.Println(myMap3)

}
