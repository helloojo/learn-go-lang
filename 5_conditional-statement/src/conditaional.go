package main

func main() {
	a := 1
	//if a {	//go if는 boolean값만
	//}
	if a == 1 {
		println("a is 1")
	} else {
		println("a is not 1")
	}
	a = 2
	if a == 1 {
		println("a is 1")
	} else if a == 2 {
		println("a is 2")
	} else {
		println("a is not 1 and 2")
	}

	if b := a * 4; b == 8 {
		println(a, " * ", 4, " = ", a*4)
	}
	//println(b);	//b는 if scope에서만 존재함
	var op = 4

	switch op {
	case 1:
		println("1")
	case 2:
		println("2")
	case 3:
		println("3")
	case 4:
		println("4")
	default:
		println("default")
	}

	switch {
	case op > 4:
		println("op > 4")
		break		//break 없어도 괜찮음
	case op < 0:
		println("op < 0")
	default:
		println("0 <= op <= 4")
	}
}
