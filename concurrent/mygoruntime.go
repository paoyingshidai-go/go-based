package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 通过异步来阻塞主线程
func Syn() {
	group := sync.WaitGroup{}
	group.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("GO2-", i)
			group.Done()
		}()
	}
	group.Wait()
}

// 通道缓存, 功能类似于 java 的 CountDownLatch
func Channel() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	//c := make(chan bool, 1)
	c := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		var value = i
		go func() {
			fmt.Println("GO1-", value)
			time.Sleep(2 * time.Second)
			c <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-c
	}
}

func Channel2() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := make(chan bool, 1)

	// 生产者
	for i := 0; i < 11; i++ {
		var value = i
		go func() {
			fmt.Println("GO1-", value)
			time.Sleep(2 * time.Second)
			fmt.Println("return-", value)
			c <- true
		}()
	}

	// 消费者
	for i := 0; i < 11; i++ {
		fmt.Println("wait")
		<-c
		time.Sleep(2 * time.Second)
	}

	fmt.Println("-------------- main ---------------")

}

// 模拟并发情况
func Channel3() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	count := 1000000
	c := make(chan bool, count)
	var value = 0

	for i := 0; i < count; i++ {
		go func() {
			value++
			c <- true
		}()
	}
	for i := 0; i < count; i++ {
		<-c
	}
	fmt.Println(value)
}

// 使用 锁 解决并发
func Channel4() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var transferLock = &sync.Mutex{}
	count := 1000000
	c := make(chan bool, count)
	var value = 0

	for i := 0; i < count; i++ {
		go func() {
			transferLock.Lock()
			defer transferLock.Unlock()
			value++
			c <- true
		}()
	}
	for i := 0; i < count; i++ {
		<-c
	}
	fmt.Println(value)
}

// 使用 channel 解决并发问题
func Channel5() {
	var value = 0
	channel := make(chan bool)

	go func() {
		for {
			<-channel
			value++
		}
	}()

	for i := 0; i < 19999; i++ {
		channel <- true
	}
	fmt.Println(value)
}

// 使用 channel 解决并发问题
func Channel6() {
	channel := make(chan int)
	var value = 0

	go func() {
		channel <- 1
		value++
	}()

	for i := 0; i < 100000; i++ {
		// var v = i
		select {
		case <-channel:
			fmt.Println(value)
		case <-time.After(2 * time.Second):
			fmt.Println("超时")
		}
	}
	fmt.Println(value)
}

func selectTest() {

	channel := make(chan int)
	var value = 0

	//go func() {
	//	for i := 0; i < 100000; i++ {
	//		go func() {
	//			channel <- 1
	//		}()
	//	}
	//}()
	for {
		select {
		case <-channel:
			value++
		case <-time.After(2 * time.Second):
			fmt.Println("超时")
		}
	}

}

func main() {

	//Channel3()
	//Channel4()
	Channel5()
	//Channel6()

	//selectTest()

}
