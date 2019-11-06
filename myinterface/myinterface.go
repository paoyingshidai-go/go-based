package myinterface

import "fmt"

type Phone interface {
	Call()
}

type Mom interface {
	Call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) Call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
	Name string
}

func (iPhone *IPhone) Call() {
	fmt.Println(iPhone.Name)
	fmt.Println("I am iPhone, I can call you!")
	iPhone.Name = "change"
}
