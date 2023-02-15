package main

import "fmt"

func foo1(a string, b int) int {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := 100
	return c
}

// 可以返回多个值 但是是匿名的
func foo2(a string, b int) (int, int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := 100
	d := 200
	return c, d
}

// 可以给返回值起名
func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println("--------------foo3--------------")
	fmt.Println("a = \n", a)
	fmt.Println("b = \n", b)

	// 给有名称的返回值变量赋值 这样就不用手动返回了
	r1 = 1000
	r2 = 2000

	return
}

// 返回值类型相同可以 省略一个
func foo4(a string, b int) (r1, r2 int) {
	fmt.Println("--------------foo4--------------")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	// r1 r2 是foo4的形参 默认初始化为0
	// 作用域在foo4
	r1 = 1000
	r2 = 2000

	return
}

func main() {
	c := foo1("hahahah", 5)
	fmt.Println("c = ", c)

	d, e := foo2("hahah", 55)
	fmt.Printf("%d %d", d, e)

	res1, res2 := foo3("hjajaj", 888)
	fmt.Println(res1)
	fmt.Println(res2)

	ret1, ret2 := foo4("hello", 789)
	fmt.Println(ret1)
	fmt.Println(ret2)
}
