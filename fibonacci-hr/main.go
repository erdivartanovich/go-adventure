package main

import "fmt"

func fibonacci(n int) int {
	x, y, z := 0, 1, 0

	f := func() {
		x, y, z = y, x+y, x
	}

	for i := 0; i <= n; i++ {
		f()
	}
	return z
}

func main() {
	var n int
	n = 6
	// fmt.Scanf("%d\n", &n)
	fmt.Println(fibonacci(n))
}
