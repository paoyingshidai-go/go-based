package main

import "fmt"

type Phone interface {
	Call()
}

type Mom interface {
	Call()
	Tell()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) Call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
	Name string
}

type Son struct {
	Name string
}

func (iPhone *IPhone) Call() {
	fmt.Println(iPhone.Name)
	fmt.Println("I am iPhone,  I can call you!")
	iPhone.Name = "change"
}

// 必须要全部关联了 Mon 才算是 Mon 的接口的实现
func (son Son) Tell() {

}
func (son Son) Call() {
	son.Name = "change"
	// 在局部作用域 Name 的值是有变更的
	fmt.Println(son.Name)
}

// 接口方法接受指针与非指针的区别，指针类可以修改接受者的值，而非指针不可以修改
// https://blog.csdn.net/idwtwt/article/details/80322714
func main() {

	IPhone := &IPhone{Name: "Michael"}
	IPhone.Call()
	fmt.Println(IPhone.Name) // 这里的name 由 Michael ——> change
	fmt.Println()

	mySon := Son{Name: "Michael"}
	mySon.Call()
	fmt.Println(mySon.Name) // 这里的值没有改变
}
