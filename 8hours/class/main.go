// @Time : 2023/02/17 22:35:05
// @Author : GuanLi
// Written by Visual Studio Code
//go:build ignore
// +build ignore

package main

import "fmt"

//如果类名首字母大写 代表其他包也可以访问 否则只能在类的内部访问
type Hero struct {
	name  string
	ad    int
	level int
}

func (self *Hero) Show() {
	fmt.Printf("self.name: %v\n", self.name)
	fmt.Printf("self.ad: %v\n", self.ad)
	fmt.Printf("self.level: %v\n", self.level)
}

// 这个this hero表示这个方法是绑定在Hero这个struct上的 this相当于一个对象
func (self *Hero) GetName() string {
	return self.name
}

func (self *Hero) SetName(newName string) {
	// self是当前调用方法的对象的一个拷贝 要加上*表示指针
	self.name = newName
}

func main() {
	hero := Hero{ //这样定义一个对象
		name:  "zhang3",
		ad:    100,
		level: 1,
	}

	hero.Show()

	fmt.Printf("hero.GetName(): %v\n", hero.GetName())

	hero.SetName("guanli")
	hero.Show()
}
