package main

import (
	"fmt"
	"reflect"
)

/*
	反射，

*/

type User struct {
	Id   int
	Name string
	Age  int
}

// 嵌套类型如何反射
type Manager struct {
	User
	title string
}

func (u User) Hello(name string) {
	fmt.Println("hello world ", name)
}

// 获取字段与方法
func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("type: ", t.Name())

	v := reflect.ValueOf(o)
	fmt.Println("Fields: ", v)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Println(f.Name, f.Type, val)
	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m.Name, m.Type)
	}

}

// 基本类型通过反射改变值
func BaseType() {
	x := 123
	v := reflect.ValueOf(&x)
	v.Elem().SetInt(999)
	fmt.Println(x)
}

// 修改自定义类型的数据
func CustomType(o interface{}) {

	v := reflect.ValueOf(o)

	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("xxx")
		return
	} else {
		v = v.Elem()
	}

	f := v.FieldByName("Name")
	if !f.IsValid() {
		fmt.Println("BAO")
		return
	}

	if f.Kind() == reflect.String {
		f.SetString("BYEBYE")
	}
}

// 通过反射调用方法
func CallMethod() {

	user := User{1, "OK", 12}

	value := reflect.ValueOf(user)
	mv := value.MethodByName("Hello")

	args := []reflect.Value{reflect.ValueOf("joe")}
	mv.Call(args)
}

func main() {

	//m := Manager{User: User{1, "OK", 12}, title: "title"}
	//t := reflect.TypeOf(m)
	//fmt.Println(t.FieldByIndex([]int{0, 0}))
	//BaseType()

	//u := User{1, "OK", 12}
	//CustomType(&u)
	//fmt.Println(u)

	//u := User{1, "OK", 12}
	//Info(u)

	//CallMethod()

}
