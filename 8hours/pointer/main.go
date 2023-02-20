//@Time : 2023/2/15 22:06
//@Author : GuanLi
//Written by Gland

package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20
	fmt.Println(a)
	fmt.Println(b)
	valueChanger(&a, &b)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(&a)

	var p *int
	p = &a

	fmt.Println(p)

	var pp **int //表示二级指针
	pp = &p
	fmt.Println(pp)
	fmt.Println(&p)
}

func valueChanger(a *int, b *int) {
	*a, *b = *b, *a
	fmt.Println("Finish!")
}
