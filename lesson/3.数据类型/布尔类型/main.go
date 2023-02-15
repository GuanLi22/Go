/*
	package main

import "fmt"

	func main() {
		var a bool = false
		var b bool = true

		var c = false
		var d = true

		e := false
		f := true

		fmt.Printf("a: %v\n", a)
		fmt.Printf("b: %v\n", b)
		fmt.Printf("c: %v\n", c)
		fmt.Printf("d: %v\n", d)
		fmt.Printf("e: %v\n", e)
		fmt.Printf("f: %v\n", f)
	}
*/

//*在if else 判断中
/*
	package main

import "fmt"

	func main() {
		age := 19
		ok := age >= 18 //ok是一个布尔类型的数据，先判断age是否大于18 是就赋值 true 否则赋值false
		fmt.Printf("%v\n", ok)
		if ok {  //判断true就执行上一个 false就执行下一个同python
			fmt.Println("你已成年")
		} else {
			fmt.Println("你还未成年")
		}
	}
*/

//*在循环中
/* package main

import "fmt"

func main() {
	count := 10
	for i := 0; i < count; i++ {
		// fmt.Printf("%v\n", i)
		fmt.Printf("i: %v\n", i)
	}
}
*/

// *在逻辑表达式中
package main

import "fmt"

func main() {
	age := 18
	gender := "男"                   //推测 &&表示python中的and 同时
	if age >= 18 && gender == "男" { //条件必须是bool类型的 不能用0或非0做判断
		fmt.Println("你是成年男子")
	}
}
