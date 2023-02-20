// @Time : 2023/02/18 14:23:45
// @Author : GuanLi
// Written by Visual Studio Code
//go:build ignore
// +build ignore

package main

//类的继承

import "fmt"

//父类
type Human struct {
	name   string
	gender string
}

func (self *Human) Eat() {
	fmt.Println("Human Eat()...")
}

func (self *Human) Walk() {
	fmt.Println("Human Walk()...")
}

//
//
//
//
// -----------------------------------
//子类

type SuperMan struct {
	Human   //表示Superman类继承了Human类的方法
	ability string
}

//重定义父类方法Eat
func (self *SuperMan) Eat() {
	fmt.Println("SuperMan Eat()...")
}

//也可以写子类的新方法
func (self *SuperMan) Fly() {
	fmt.Println("SuperMan Fly()...")
}

func (self *SuperMan) PrintInformation() {
	fmt.Println(self.name)
	fmt.Println(self.gender)
	fmt.Println(self.ability)
	fmt.Println(self.Human)
}

func main() {
	h := Human{"guanli", "man"}
	h.Eat()
	h.Walk()
	fmt.Println("-----------------------")

	//定义一个子类对象
	s := SuperMan{Human{"Alex", "man"}, "fly"} //继承哪个要在里面这样写

	//如果觉得别扭可以这样写
	var s1 SuperMan
	s1.name = "Alex"
	s1.gender = "man"
	s1.ability = "fly"

	s.Walk() //父类 方法
	s.Eat()  //子类 重写之后的方法
	s.Fly()  //子类 新方法
	fmt.Println("-----------------------")

	// 这样定义后是和以前一样的
	s1.Walk()
	s1.Eat()
	s1.Fly()
	fmt.Println("-----------------------")

	s1.PrintInformation()
}
