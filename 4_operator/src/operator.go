package main

func main() {
	var a, b int = 10, 20
	println(a, " + ", b, " = ", a+b)
	println(a, " == ", b, " = ", a == b)
	println(a, " != ", b, " = ", a != b)
	println(a, " != ", b, " && ", a, " == ", b, " = ", a != b && a == b)
	println(a, " != ", b, " || ", a, " == ", b, " = ", a != b || a == b)
	var c = &a;
	println(c);
	println(*c)
	// c++
	*c++
	println(*c, a)
}
