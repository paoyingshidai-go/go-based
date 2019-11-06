package main

import (
	"fmt"
	"workspace2/myinterface"
)

func main() {

	//_,numb,strs := function.Numbers() //只获取函数返回值的后两个
	//fmt.Println(numb,strs)

	// 数组
	//array.PrintArray()

	// 指针
	//point.MyPoint()

	// 遍历范围
	//_range.RangeTest()

	// struct
	//mystruct.PrintBook()

	// 设置实体类
	//mystruct.SetBookProperties()

	// 设置和输出 map 内容
	//mymap.SetAndPrintMap()

	// Map 测试
	//mymap.MapDelete()

	// 递归， 斐波那契
	//fmt.Println(recursion.Factorial(12))

	// 类型转换
	//typechange.TypeChange()

	//go concurrent.Sleep("world")
	//concurrent.Sleep("hello")

	//concurrent.TestSum()

	//concurrent.TestClose()

	// error

	// 正常情况
	//if result, errorMsg := myerror.Divide(100, 10); errorMsg == "" {
	//	fmt.Println("100/10 = ", result)
	//}
	//// 当被除数为零的时候会返回错误信息
	//if _, errorMsg := myerror.Divide(100, 0); errorMsg != "" {
	//	fmt.Println("errorMsg is: ", errorMsg)
	//}

	// 接口测试
	//phone := myinterface.IPhone{Name: "IPhone"}
	//phone.Call()
	//TypeCheck(phone)
	//fmt.Println(phone.Name)
	//
	//nokia := myinterface.NokiaPhone{}
	//nokia.Call()
	//TypeCheck(nokia)

	// 闭包
	//n := 0
	//f := func() int {
	//	n += 1
	//	return n
	//}
	//fmt.Println(f())

	// 函數作為返回值
	//method := function.Increase()
	//fmt.Println(method())

	//adder := function.Adder()
	//fmt.Println(adder(2))
	//fmt.Println(adder(2))
	//fmt.Println(adder(2))

	//me := FuncAsParamImpl
	//function.FuncAsParam(me)

	// goruntimes

	c := make(chan bool)
	go func() {
		fmt.Println("GO, Go, GO")
		c <- true
		close(c) // 关闭通道
	}()
	// 这里等待 go 函数执行完， 相当于 java 里面的 join
	//<- c
	fmt.Println("GO, Go")

	// 迭代通道的时候必须先关闭通道
	for v := range c {
		println(v)
	}
	fmt.Println("GO")
}

// 这个是函数的模板的实现
func FuncAsParamImpl(in string) string {
	fmt.Println("方法的实现, 入参： ", in)
	return in
}

// 接口类型判断
func TypeCheck(phoneType interface{}) {
	switch v := phoneType.(type) {
	case myinterface.IPhone:
		fmt.Println("IPhone type:", v)
	case myinterface.NokiaPhone:
		fmt.Println("NokiaPhone type:")
	default:
		fmt.Println("unknowde type")
	}
}
