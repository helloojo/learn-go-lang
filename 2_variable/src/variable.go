package main

import "fmt"

func main() {
	var a int
	fmt.Println(a) //default 0

	var aa int64 = 1
	fmt.Println(aa)

	var f float32 = 1.5
	fmt.Println(f)

	var d float64 = 1.5
	fmt.Println(d)

	var i, j, k int = 1, 2, 3
	fmt.Println(i, j, k)

	const c int = 10
	const s string = "Hello"
	fmt.Println(c, s)

	const (
		qwerty         = "qwerty"
		asdfgh         = "asdfgh"
		zxcv   int8    = 15
		flt    float32 = 11.
	)
	fmt.Println(qwerty, asdfgh, zxcv, flt)

	const (
		A = iota // 0, 1, 2 assign
		B
		C
	)

	fmt.Println(A, B, C)
}
