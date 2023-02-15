/*
package main

import "fmt"

	func main() {
		name := `
			aaaaaaaaaaaaa
			aaaaaaaaaaaaaaaaaaaa
			aaaaaaaaaaaaaaaaa
			aaaaaaaaaaaaaaaaa
		`  //多行字符串用``反引号
		var myname string = "guanli"
		var name1 = "guanli1"
		name2 := "guanli2"
		fmt.Printf("%v\n", name)
		fmt.Printf("%v\n", myname)
		fmt.Printf("%v\n", name1)
		fmt.Printf("%v\n", name2)
	}
*/
package main

import (
	"bytes"
	"fmt"
	"strings"
)

// *字符串拼接
func main() {
	name := "guanli"
	age := "15"
	msg := name + " " + age //msg := name + age也可以这里是
	fmt.Printf("msg:%v\n", msg)

	msg = ""
	msg += name
	msg += " "
	msg += age
	fmt.Printf("msg:%v\n", msg)

	//调用fmt的这个sprintf功能 对字符串进行格式化 不会直接打印 会返回个值
	msg = fmt.Sprintf("msg:%s %s", name, age)
	fmt.Printf("%v\n", msg)

	//还没学
	msg = ""
	msg = strings.Join([]string{name, age}, " ")
	fmt.Printf("msg:%v\n", msg)

	//还没学
	var buffer bytes.Buffer
	buffer.WriteString("guanli")
	buffer.WriteString(" 15")
	fmt.Printf("buffer.String(): %v\n", buffer.String())
	fmt.Println("---------------------")

	zhuanyi()
}

// *转义字符
func zhuanyi() {
	// \n 换行
	s := "hello"
	print(s + "\n")
	print(s)
	s = "hello \n world" //可以写在这里
	fmt.Printf("s: %v\n", s)

	// \t相当于按了tab键
	s = "hello\tworld"
}
