package main

import "fmt"

func solveMeFirst(a, b uint) uint {
	return a + b
}

func main() {
	var a, b, res uint

	fmt.Scan(&a, &b)

	res = solveMeFirst(a, b)
	fmt.Println(res)
}
