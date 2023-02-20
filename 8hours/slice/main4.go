// @Time : 2023/02/16 22:47:50
// @Author : GuanLi
// Written by Visual Studio Code

package main

import "fmt"

//切片的截取

func main() {
	s := []int{1, 2, 3} // len = 3, cap = 3
	// s = [1,2,3]

	s1 := s[0:2] // 左含右不含 表示的是[1,2]

	fmt.Println(s1)

	s1[0] = 100 //虽然只是截取的部分 但是他们指向的内存区域是一样的

	fmt.Println(s)
	fmt.Println(s1)

	//copy函数 可以将底层的数组slice一起进行拷贝(也就是深拷贝)

	s2 := make([]int, 3) //此时s2 = [0,0,0]

	//将s中的值依次拷贝到s2中
	copy(s2, s)
	fmt.Println(s2) // 现在s2的值和s相等但是 内存地址不同

}
