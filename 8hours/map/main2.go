// @Time : 2023/02/17 19:46:38
// @Author : GuanLi
// Written by Visual Studio Code

package main

import "fmt"

//map的使用

func printMap(cityMap map[string]string) {
	//这里的cityMap 是一个引用传递
	for key, value := range cityMap {
		fmt.Printf("key: %v\n", key)
		fmt.Printf("value: %v\n", value)
	}
}

func changeValue(cityMap map[string]string) {
	cityMap["England"] = "London" //也就是说在函数中修改会生效
	for key, value := range cityMap {
		fmt.Printf("key: %v\n", key)
		fmt.Printf("value: %v\n", value)
	}
}

func main() {
	cityMap := make(map[string]string)

	// 添加
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"

	// 遍历
	for key, value := range cityMap {
		fmt.Printf("key: %v\n", key)
		fmt.Printf("value: %v\n", value)
	}
	fmt.Println("---------------------------------")

	//修改
	cityMap["USA"] = "DC"
	for key, value := range cityMap {
		fmt.Printf("key: %v\n", key)
		fmt.Printf("value: %v\n", value)
	}
	fmt.Println("---------------------------------")

	//删除
	//使用delete函数
	delete(cityMap, "China") //前面写容器 后面写删除行对应的key
	for key, value := range cityMap {
		fmt.Printf("key: %v\n", key)
		fmt.Printf("value: %v\n", value)
	}
	fmt.Println("---------------------------------")

	printMap(cityMap)

	fmt.Println("---------------------------------")

	changeValue(cityMap)
}
