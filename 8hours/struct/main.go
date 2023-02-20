// @Time : 2023/02/17 21:37:59
// @Author : GuanLi
// Written by Visual Studio Code
//go:build ignore
// +build ignore

package main

import "fmt"

//声明一种新的数据类型myInt 实际上是int的别名
type myInt int

type book struct {
	title  string
	author string
}

func changeBook(book book) {
	//传递一个book的副本
	book.author = "666" //这样修改不了
}

func changeBook2(book *book) {
	//指针传递
	book.author = "666" //这样就可以修改
}

func main() {
	/*var a myInt = 10
	fmt.Printf("a: %v\n", a)
	fmt.Printf("type of a: %T\n", a)*/

	var book1 book //相当于新建了一个对象
	book1.title = "golang"
	book1.author = "zhang3"

	fmt.Printf("%v\n", book1)

	changeBook(book1)

	fmt.Printf("%v\n", book1)

	changeBook2(&book1)

	fmt.Printf("%v\n", book1)
}
