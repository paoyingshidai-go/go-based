package main

import (
	"fmt"
	"os"
	"time"
)

// go 的异常处理

func Defer() {
	defer fmt.Println("defer main") // will this be called when panic?
	var user = os.Getenv("USER_")
	go func() {
		defer func() {
			fmt.Println("defer caller")
			if err := recover(); err != nil {
				fmt.Println("recover success.")
			}
		}()
		func() {
			defer func() {
				fmt.Println("defer here")
			}()

			if user == "" {
				panic("should set user env.")
			}
			fmt.Println("after panic")
		}()
	}()

	time.Sleep(1 * time.Second)
	fmt.Printf("get result %d\r\n", 1)
}

// https://codeday.me/bug/20190910/1801997.html
/**
在循环中处理异常需要将整个模块封装成函数
*/
func ExceptionInLoop() {
	//channel := make(chan bool)

	go func() {
		// 1 在这里需要你写算法
		// 2 要求每秒钟调用一次proc函数
		// 3 要求程序不能退出

		for {
			func() {
				defer func() {
					if err := recover(); err != nil {
						fmt.Println("recover success.")
					}
				}()
				time.Sleep(1 * time.Second)
				panic("ok")
			}()
		}
		//channel <- true
	}()
	select {
	//case <-channel:
	//	fmt.Println("----")
	//case <-time.After(100 * time.Second):
	//	fmt.Println("time out")
	}
}

func main() {

	ExceptionInLoop()
}
