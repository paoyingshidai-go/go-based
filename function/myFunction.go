/*

 */
package function

import "fmt"

//一个可以返回多个值的函数
func Numbers() (int, int, string) {
	a, b, c := 1, 2, "str"
	return a, b, c
}

// 闭包
func Increase() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

// 闭包，x 的生命周期是一直伴随着 Adder
func Adder() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}

// 函数作为参数传入
type FuncParam func(in string) string

// 这种写法跟 java 的函数式编程一样，go 可以自定义入参的格式，java 可以使用 @FunctionalInterface， 然后使用 lambda 实现约定的入参格式
func FuncAsParam(fun FuncParam) {
	fmt.Println("开始调用传入的参数函数")
	str := "hello"
	fun(str)
}
