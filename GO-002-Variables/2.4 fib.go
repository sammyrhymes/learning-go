package main

import "fmt"

// Fibonacci sequence 0,1,1,2,3,5,8

func fib(n int) (fibNo int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
func main() {
	fmt.Println(fib(4))
}
