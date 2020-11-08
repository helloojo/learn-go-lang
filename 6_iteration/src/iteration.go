package main

func main() {
	var sum int = 0

	for i := 1; i <= 10; i++ {
		sum += i
	}
	println(sum)

	is_ten := false
	a := 0
	for is_ten {
		a++
		if a == 10 {
			is_ten = true
		}
	}
	println("a = ", a, is_ten)

	arrs := []int{1, 2, 3}

	for idx, num := range arrs {
		println("arrs[", idx, "]", " = ", num)
	}
}
