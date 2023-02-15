package main

import (
	"fmt"
)

const age_float float64 = 12.0 //*类型也可以省略
const age = 12                 //*必须赋值
const (
	height = 114
	width  = 514
)
const a, b, c = 1, 2, 3
const (
	d = "d"
	f
	e
)

func main() {
	const (
		a1 = iota //*iota 对变量赋值，原始值是0,每赋一次+1
		a2        //1
		a3        //2
	)
	const (
		a11 = iota //0
		_          //*跳过一个1
		a22 = iota //2
	)
	const (
		a111 = iota //0
		a222 = 100  //*中间插值会自动跳过
		a333 = iota //2
	)

	fmt.Printf("age_float: %v\n", age_float)
	fmt.Printf("age: %v\n", age)
	fmt.Printf("height: %v\n", height)
	fmt.Printf("width: %v\n", width)
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)
	fmt.Printf("d: %v\n", d)
	fmt.Printf("f: %v\n", f)
	fmt.Printf("e: %v\n", e)
	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)
	fmt.Printf("a3: %v\n", a3)
	fmt.Printf("a11: %v\n", a11)
	fmt.Printf("a22: %v\n", a22)
	fmt.Printf("a111: %v\n", a111)
	fmt.Printf("a222: %v\n", a222)
	fmt.Printf("a333: %v\n", a333)
}
