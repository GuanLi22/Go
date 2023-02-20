package main

import (
	_ "Go/8hours/init/lib1"
	. "Go/8hours/init/lib2" // 这样会把整个包里面的内容全导入进来 之后 可以直接调用(不建议)
	myLib2 "Go/8hours/init/lib2"
	//"Go/8hours/init/lib2"
	//"Go/8hours/init/lib1"
)

// 正常来说 导入包但不使用会报错 但是如果我不想使用里面的方法 只想执行init()方法
// 可以使用_来进行匿名导入

// 还可以给函数起别名 这样导入时就可以用别名导入

func main() {
	//lib1.Lib1Test()
	myLib2.Lib2Test()
	Lib2Test() // 这样全导入进来之后 可以直接调用
}
