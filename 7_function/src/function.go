package main

func main() {
	func1()
	func2(10)
	func3("string")
	n := 5
	func_reference(&n)
	println("n = ", n)
	func_variadic("a")
	func_variadic("a", "b")
	func_variadic("a", "ab", "abc", "abcd")
	avg := average(1, 2, 3, 4, 5)
	println("average(1,2,3,4,5) = ", avg)

	sum, length := get_sum_and_length(1, 2, 3, 4, 5)
	println("sum(1,2,3,4,5) =", sum, ", length: ", length)
	sum2, length2 := get_sum_and_length2(1, 2, 3, 4, 5)
	println("sum2(1,2,3,4,5) =", sum2, ", length: ", length2)
}

func func1() {
	println("func1()")
}

func func2(num int) {
	println("func2(", num, ")")
}

func func3(str string) {
	println("func3(", str, ")")
}

func func_reference(num *int) {
	println("func_reference(", num, ")")
	*num = 15
}

func func_variadic(strs ...string) {
	println("parameter length: ", len(strs))
	for idx, str := range strs {
		println("strs[", idx, "]", "=", str)
	}
}

func average(nums ...int) float64 {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return float64(sum / len(nums))
}

func get_sum_and_length(nums ...int) (int, int) {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum, len(nums)
}

func get_sum_and_length2(nums ...int) (sum int, length int) {
	for _, num := range nums {
		sum += num
	}
	length = len(nums)
	return
}
