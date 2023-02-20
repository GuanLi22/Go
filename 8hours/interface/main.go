// @Time : 2023/02/19 21:17:20
// @Author : GuanLi
// Written by Visual Studio Code

package main

import "fmt"

// 本质是一个指针
type AnimalIF interface {
	Sleep()
	Getcolor() string //获取动物的颜色
	GetType() string  //获取动物的种类
}

// 定义一个具体的类
type Cat struct {
	color string //猫的颜色
}

func (self *Cat) Sleep() {
	fmt.Println("Cat Sleep()...")
}

func (self *Cat) Getcolor() string {
	return self.color
}

func (self *Cat) GetType() string {
	return "Cat"
}

// 具体的类
type Dog struct {
	color string
}

func (self *Dog) Sleep() {
	fmt.Println("Dog Sleep()...")
}

func (self *Dog) Getcolor() string {
	return self.color
}

func (self *Dog) GetType() string {
	return "Dog"
}

func showAnimal() {

}

func main() {
	var animal AnimalIF //接口的数据类型 父类的指针
	animal = &Cat{"Green"}

	animal.Sleep() //调用的就是Cat的Sleep()方法

	animal = &Dog{"Yellow"}

	animal.Sleep() //调用了Dog的Sleep()方法 多态的现象
}
