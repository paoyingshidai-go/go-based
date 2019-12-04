/*

 */
package main

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

// 函数的实现, 执行方法进行拦截
func (f FuncParam) Call(w string) string {
	fmt.Println("Call")
	return f(w)
}

//
func FuncCall(f FuncParam) {
	f.Call("Michael")
}

// 这个是函数的模板的实现
func FuncAsParamImpl(in string) string {
	fmt.Println("方法的实现, 入参： ", in)
	return in
}

// 这种写法跟 java 的函数式编程一样，go 可以自定义入参的格式，java 可以使用 @FunctionalInterface， 然后使用 lambda 实现约定的入参格式
func FuncAsParam(fun FuncParam) {
	fmt.Println("开始调用传入的参数函数")
	str := "hello"
	fun(str)
}

// 也可以不需要方法声明1约束
func FuncAsParam2(fun func(string) string) {
	fmt.Println("开始调用传入的参数函数")
	str := "hello"
	fun(str)
}

func main() {

	// 函數作為返回值
	//method := Increase()
	//fmt.Println(method())
	//
	//adder := Adder()
	//fmt.Println(adder(2))
	//fmt.Println(adder(2))
	//fmt.Println(adder(2))

	me := FuncAsParamImpl
	//FuncAsParam(me)
	FuncCall(me)
}
