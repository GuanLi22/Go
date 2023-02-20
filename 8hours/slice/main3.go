//@Time : 2023/2/16 16:28
//@Author : GuanLi
//Written by Gland

/*切片的追加*/

package main

import "fmt"

func main() {
	var numbers = make([]int, 3, 5) //这个切片的长度为3 内存中开辟的长度为5

	fmt.Printf("len = %d, cap = %d, slice = %v\n ", len(numbers), cap(numbers), numbers)

	numbers = append(numbers, 1) // 向numbers切片追加元素1 前提是要小于等于内存中开辟的长度
	// 现在 numbers len = 4, [0,0,0,1], 容量还是5
	fmt.Printf("len = %d, cap = %d, slice = %v\n ", len(numbers), cap(numbers), numbers)

	numbers = append(numbers, 2)
	//现在 numbers len = 5  [0,0,0,1,2]  cap = 5
	fmt.Printf("len = %d, cap = %d, slice = %v\n ", len(numbers), cap(numbers), numbers)

	//向一个cap 满了的slice 追加元素
	numbers = append(numbers, 3)
	//这时go编辑器会帮助你再开辟一个内存 并且长度为原来内存的2倍 然后把这个元素追加进去
	//现在 numbers len = 5  [0,0,0,1,2,3]  cap = 10
	fmt.Printf("len = %d, cap = %d, slice = %v\n ", len(numbers), cap(numbers), numbers)

	var numbers2 = make([]int, 3)
	fmt.Printf("len = %d, cap = %d, slice = %v\n ", len(numbers2), cap(numbers2), numbers2)
	numbers2 = append(numbers2, 1) //长度还是为原来内存的2倍
	fmt.Printf("len = %d, cap = %d, slice = %v\n ", len(numbers2), cap(numbers2), numbers2)
}
