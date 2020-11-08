package main

import "fmt"

func main() {
	var slice1 []int
	slice1 = []int{0, 1, 2}
	println(len(slice1), cap(slice1))
	slice2 := make([]int, 10, 20) //length 10, capacity 20
	println(len(slice2), cap(slice2))
	for i := 0; i < len(slice2); i++ {
		slice2[i] = i
	}
	println("---------------------")
	slice3 := slice2[3:8] // [3,8)
	fmt.Println("slice3: ", slice3)
	slice4 := slice2[5:] // [5,10)
	fmt.Println("slice4: ", slice4)
	slice5 := append(slice3, 10)
	fmt.Println("slice5: ", slice5) // slice3, 10
	slice6 := append(slice5, 11, 12, 13)
	fmt.Println("slice6: ", slice6) // slice6, 11, 12, 13
	println("---------------------")
	slice2[4] = 12 // 깊은 복사가 아님
	fmt.Println("slice3: ", slice3)
	fmt.Println("slice5: ", slice5)
	fmt.Println("slice6: ", slice6)
	println("---------------------")
	slice7 := append(slice3, 20, 21, 22)
	fmt.Println("slice5: ", slice5) // 깊은 복사가 아니기때문에 10, 11, 12 덮어씌워짐!
	fmt.Println("slice6: ", slice6)
	fmt.Println("slice7: ", slice7)
	println("---------------------")
	slice8 := make([]int, len(slice7), cap(slice7))
	copy(slice8, slice7)
	slice8[2] = 10
	fmt.Println("slice7: ", slice7)
	fmt.Println("slice8: ", slice8)
	println("---------------------")
	slice3 = append(slice3, slice4...)
	fmt.Println("slice3: ", slice3)
	slice4[2] = 999
	fmt.Println("slice3: ", slice3)
	fmt.Println("slice4: ", slice4)
	println("---------------------")
	println("slice2 len, cap:", len(slice2), cap(slice2))
	diff := cap(slice2) - len(slice2)
	for i := 0; i < diff; i++ {
		slice2 = append(slice2, i)
	}
	fmt.Println("slice2 len, cap, val:", len(slice2), cap(slice2), slice2)
	slice2 = append(slice2, 999)
	fmt.Println("slice2 len, cap, val:", len(slice2), cap(slice2), slice2)
	fmt.Println("slice3 len, cap, val:", len(slice3), cap(slice3), slice3) //이제 slice2, slice3은 별도의 메모리공간
	/**
	slice는 포인터를 이용해 연속된 메모리공간을 할당한다.
	append로 새로운 값을 추가해도 같은 메모리공간을 공유한다.
		len < cap => 같은 메모리 공간에 값 append
		len >= cap => 새로운 메모리공간(cap*2)에 값 복사, append
	깊은 복사를 하려면 copy 메서드를 사용해야한다.
	*/
}
