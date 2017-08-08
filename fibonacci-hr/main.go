package main

import "fmt"

func fib(n int) int {
	if n == 1 {
		return 1
	}
	if n == 0 {
		return 0
	}

	return fib(n-2) + fib(n-1)
}

func main() {
	var n int
	n = 6
	// fmt.Scanf("%d\n", &n)
	fmt.Println(fib(n))
}
