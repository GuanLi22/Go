// @Time : 2023/2/1 9:43
// @Author : GuanLi
// Written by Gland
package main

import (
	"fmt"
	"strings"
)

var s = "string"

// 转义符
/*

\r 表示回到行首 如果后面还有字符会覆写
\n 表示换行
\t 制表符 = tab
\' 单引号
\" 双引号
\\ 反斜杠

*/

// 多行字符串
var s2 = `
	hahah
	6666
			777`

func main() {
	fmt.Println("str := \"C:\\Users\\FAQ11\\Desktop\\codes\\code\\Go\"")
	fmt.Println("我是前面的在这之后的不会被覆写\r前面被覆写")
	fmt.Println("Hello\nworld")
	fmt.Println("Hello\tworld")

	fmt.Println(s2)

	/*字符串的常用操作*/
	// 求长度
	fmt.Println(len(s))

	// 字符串拼接
	a := "hello"
	b := "world"
	fmt.Println(a + b)

	fmt.Printf("%s%s\n", a, b)

	res := fmt.Sprintf("%s%s\n", a, b) // 这个需要把输出的值存起来
	fmt.Println(res)

	//字符串分割
	dir := `C:\Users\FAQ11\Desktop\codes\code\Go`
	ret := strings.Split(dir, "\\")
	fmt.Println(ret) // 分割后类似python中的列表 可以取值

	// 判断是否包含
	fmt.Println(strings.Contains(res, "hello"))
	// 判断前后缀
	fmt.Println(strings.HasSuffix(res, "world\n")) // 以world为后缀
	fmt.Println(strings.HasPrefix(res, "hello"))   // 以hello为前缀

	//字符串中字符出现的位置
	ss := "abcdeb"
	fmt.Println(strings.Index(ss, "b"))     //返回第一个
	fmt.Println(strings.LastIndex(ss, "b")) //返回最后一个

	//切片拼成字符串
	fmt.Println(strings.Join(ret, "\\"))
}
