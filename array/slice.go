package main

import "fmt"

func SliceCopy() {
	s := []int{1, 2, 3}
	fmt.Println(s) //[1 2 3]
	copy(s, []int{4, 5, 6, 7, 8, 9})
	fmt.Println(s) //[4 5 6]

	s2 := []int{4, 5, 6, 7, 8, 9}
	copy(s2, []int{1, 2, 3})
	fmt.Println(s2)
}

// 切片元素的添加，扩容
// https://blog.csdn.net/u013474436/article/details/88770501
func AppendSlice() {

	s := []int{1, 2, 3}
	fmt.Println(s) //[1 2 3]

}

func main() {

	SliceCopy()

}
