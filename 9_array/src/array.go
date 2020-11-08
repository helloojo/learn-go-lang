package main

import "fmt"

func main() {
	var arr1 [3]int
	arr1[0] = 0
	arr1[1] = 1
	arr1[2] = 2
	fmt.Println(arr1)
	println(arr1[0], arr1[1], arr1[2])
	var arr2 = [2]int{0, 1}
	var arr3 = [...]int{0, 1, 2}
	println(arr2[0], arr2[1])
	println(arr3[0], arr3[1], arr3[2])

	var arr4 [2][2]int
	arr4[1][1] = 1
	println(arr4[1][1]) // 1

	var arr5 = [2][3]int{
		{0, 1, 2},
		{3, 4, 5},
	}
	println(arr5[1][2]) // 5

	//var arr6 = [2][...]int{
	//	{0, 1},
	//	{2, 3, 4},
	//}		//다차원 가변배열 불가능
}
