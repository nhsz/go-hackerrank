package main

import "fmt"

func main() {
	var a0, a1, a2 int
	var b0, b1, b2 int
	fmt.Scan(&a0, &a1, &a2)
	fmt.Scan(&b0, &b1, &b2)

	ar0, br0 := comparePoints(a0, b0)
	ar1, br1 := comparePoints(a1, b1)
	ar2, br2 := comparePoints(a2, b2)

	alicePoints := ar0 + ar1 + ar2
	bobPoints := br0 + br1 + br2

	fmt.Println(alicePoints, bobPoints)
}

func comparePoints(ai, bi int) (aRes, bRes int) {
	if ai > bi {
		return 1, 0
	} else if bi > ai {
		return 0, 1
	} else {
		return 0, 0
	}
}
