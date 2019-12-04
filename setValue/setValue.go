/*
	值传递，引用传递
*/
package main

import "fmt"

/* 值传递*/
/* 定义相互交换值的函数 */
func Swap(x, y int) int {
	var temp int

	temp = x /* 保存 x 的值 */
	x = y    /* 将 y 值赋给 x */
	y = temp /* 将 temp 值赋给 y*/

	return temp
}

/* 引用传递*/
/* 定义交换值函数*/
func Swap2(x *int, y *int) {
	var temp int
	temp = *x /* 保持 x 地址上的值 */
	*x = *y   /* 将 y 值赋给 x */
	*y = temp /* 将 temp 值赋给 y */
}

func main() {
	x := 1
	y := 2

	//Swap(x, y)
	//Swap2(x, y)
	Swap2(&x, &y)
	fmt.Printf("x = %d, y = %d", x, y)

}
