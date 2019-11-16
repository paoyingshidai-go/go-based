package main

import (
	"container/list"
	"fmt"
)

type Book struct {
	Name string
}

func PrintArray() {
	var n [10]int /* n 是一个长度为 10 的数组 */
	var i, j int

	/* 为数组 n 初始化元素 */
	for i = 0; i < 10; i++ {
		n[i] = i + 100 /* 设置元素为 i + 100 */
	}

	/* 输出每个数组元素的值 */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
}

func PrintBooks() {

	var books [10]Book

	var i, j int
	for i = 0; i < 10; i++ {
		books[i] = Book{"michael"}
	}

	for j = 0; j < 10; j++ {
		fmt.Println(books[j].Name)
	}

}

func PrintList() {
	list := list.New()

	list.PushBack(Book{Name: "水浒传"})
	list.PushBack(Book{Name: "水浒传2"})

	for e := list.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func main() {

	PrintBooks()

}
