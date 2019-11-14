package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	//ValueOrPoint()
	c := make(map[string]string)
	c["1"] = "hello"
	ValueOrPoint2(c)
	for k, v := range c {
		fmt.Println(k, ":", v)
	}
	//ThreadSafe2()

}

// map 是线程不安全的
func ThreadSafe() {
	//c := make(map[string]int)
	c := make(map[string]int)

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				c[fmt.Sprintf("%d", j)] = j
			}
		}()
	}
	time.Sleep(3 * time.Second)
	fmt.Println(c)
}

// https://blog.csdn.net/qq_39920531/article/details/88062841
func ThreadSafe2() {
	//开箱即用
	var sm sync.Map
	//store 方法,添加元素
	sm.Store(1, "a")
	//Load 方法，获得value
	if v, ok := sm.Load(1); ok {
		fmt.Println(v)
	}
	//LoadOrStore方法，获取或者保存
	//参数是一对key：value，如果该key存在且没有被标记删除则返回原先的value（不更新）和true；不存在则store，返回该value 和false
	if vv, ok := sm.LoadOrStore(1, "c"); ok {
		fmt.Println(vv)
	}
	if vv, ok := sm.LoadOrStore(2, "c"); !ok {
		fmt.Println(vv)
	}
	//遍历该map，参数是个函数，该函数参的两个参数是遍历获得的key和value，返回一个bool值，当返回false时，遍历立刻结束。
	sm.Range(func(k, v interface{}) bool {
		fmt.Print(k)
		fmt.Print(":")
		fmt.Print(v)
		fmt.Println()
		return true
	})
}

// map 是值传递
func ValueOrPoint2(param map[string]string) {
	param["1"] = "Michael"
}

// 值传递还是引用传递
func ValueOrPoint() {
	m := map[string]int{"value": 0}
	m1 := m
	fmt.Println("m =", m)
	fmt.Println("m1 =", m1)
	m1["value"] = 1
	fmt.Println("m =", m)
	fmt.Println("m1 =", m1)
}

func SetAndPrintMap() {
	var countryCapitalMap map[string]string /*创建集合 */
	countryCapitalMap = make(map[string]string)

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India "] = "新德里"

	/*使用键输出地图值 */
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	/*查看元素在集合中是否存在 */
	capital, ok := countryCapitalMap["American"] /*如果确定是真实的,则存在,否则不存在 */
	/*fmt.Println(capital) */
	/*fmt.Println(ok) */
	if ok {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}
}

func MapDelete() {
	/* 创建map */
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}

	fmt.Println("原始地图")

	/* 打印地图 */
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	/*删除元素*/
	delete(countryCapitalMap, "France")
	fmt.Println("法国条目被删除")

	fmt.Println("删除元素后地图")

	/*打印地图*/
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}
}
