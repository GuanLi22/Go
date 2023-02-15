package main

import "fmt"

func k() { //函数类型

}

func main() {
	var (
		name = "tom"
		age  = 20
		b    = true
	)
	a := 100 //通道类型
	p := &a

	f := [2]int{1, 2} //数组类型

	g := []int{1, 2, 3} //切片类型

	fmt.Printf("%T\n", name) //%表示数据类型
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", p)
	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", g)
	fmt.Printf("%T\n", k)
}
