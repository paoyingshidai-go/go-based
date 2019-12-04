package main

/*
	斐波那契数列
*/
func Factorial(n uint64) (result uint64) {
	if n == 1 || n == 2 {
		return 1
	} else {
		return Factorial(n-2) + Factorial(n-1)
	}
}

func main() {
	println(Factorial(4))
}
