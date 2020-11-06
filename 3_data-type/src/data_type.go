package main

import "fmt"

func main() {
	var f bool = false
	fmt.Println(f)

	a := 10 // a variable create
	fmt.Println(a)
	a = 15
	fmt.Println(a)

	rawLiteral := `test\n
  test\n
                                ererewrwrwe`
	interLiteral := "abcabc\n" +
		"abcd" +
		"               avbcc"
	fmt.Println(rawLiteral)
	fmt.Println(interLiteral)

	var i int = 100
	//	var u uint = i	implicit conversion X
	var u uint = uint(i)
	var d float64 = float64(u)
	//	var t bool = bool(i)	int -> bool conversion X
	//	var t int = int(f)		bool -> int conversion X
	fmt.Println(i, u, d)

	var str string = "ABC"
	bytes := []byte(str)
	str2 := string(bytes)
	fmt.Println(str, bytes, str2)
}
