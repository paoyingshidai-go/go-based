package main

import "fmt"

func SliceCopy() {
	s := []int{1, 2, 3}
	fmt.Println(s) //[1 2 3]
	copy(s, []int{4, 5, 6, 7, 8, 9})
	fmt.Println(s) //[4 5 6]

	s2 := []int{4, 5, 6, 7, 8, 9}
	copy(s2, []int{1, 2, 3})
	fmt.Println(s2)
}

func main() {

	SliceCopy()

}
