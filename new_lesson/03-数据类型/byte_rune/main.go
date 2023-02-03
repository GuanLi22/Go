//@Time : 2023/2/2 17:07
//@Author : GuanLi
//Written by Gland

package main

import (
	"fmt"
)

func main() {
	s := "Hello长春"
	n := len(s)    // len求的是byte字节的数量
	fmt.Println(n) // 求s的长度

	for _, c := range s { // 从字符串中拿出具体字符
		fmt.Printf("%c\n", c)
	}

	// 修改字符串
	// 中文
	s1 := "白萝卜"
	s2 := []rune(s1)        //把字符串强制转成了一个rune切片
	s2[0] = '红'             // 切片里的每个元素都是字符 所以要修改也要用字符改
	fmt.Println(string(s2)) // 把rune切片转成字符串
	// 英文
	a1 := "hello"
	a2 := []rune(a1)
	a2[1] = 'a'
	fmt.Println(string(a2))

	c := "a" // string
	d := 'a' // rune = int32
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", d)

	e := "H"
	f := 'H'
	fmt.Printf("%T\n", e)
	fmt.Printf("%T\n", f)
	fmt.Printf("%d\n", f)

	/*类型转换*/
	n2 := 10
	nf := float64(n2)
	fmt.Printf("%T\n", nf)
}
