package main

func main() {
	sum := func(nums []int) int {
		s := 0
		for _, n := range nums {
			s += n
		}
		return s
	}
	average := func(nums ...int) float32 {
		s := sum(nums)
		return float32(s / len(nums))
	}
	println("average(1,2,3,4,5) =", average(1, 2, 3, 4, 5))

	println("avg(sum,1,2,3,4,5) =", avg(sum, 1, 2, 3, 4, 5))

	println(calc_result(func(a int, b int) int {
		return a + b
	}, 2, 2))

	println(calc_result(func(a int, b int) int {
		return a - b
	}, 2, 2))

	println(calc_result(func(a int, b int) int {
		return a * b
	}, 2, 2))

	println(calc_result(func(a int, b int) int {
		return a / b
	}, 2, 2))

	println(calc_result(func(a int, b int) int {
		return a % b
	}, 2, 2))

	sum_cont1 := sum_container()
	println(sum_cont1(5))   //	5
	println(sum_cont1(10))  // 15
	println(sum_cont1(100)) // 115

	sum_cont2 := sum_container()
	println(sum_cont2(2))  // 2
	println(sum_cont2(12)) // 14
}

func avg(sum func(nums []int) int, n ...int) float32 {
	return float32(sum(n) / len(n))
}

type calc func(int, int) int

func calc_result(f calc, a int, b int) int {
	return f(a, b)
}

func sum_container() func(int) int {
	sum := 0
	return func(num int) int {
		sum += num
		return sum
	}
}
