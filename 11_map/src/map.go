package main

import "fmt"

func main() {
	var intStrMap map[int]string // nil
	fmt.Println(intStrMap)
	intStrMap = make(map[int]string) // map init
	fmt.Println(intStrMap)
	intStrMap[10] = "abc"
	intStrMap[20] = "bbc"

	mapInit := map[string]string{
		"t": "ttt",
		"a": "aaa",
		"b": "bbb",
	}
	fmt.Println(mapInit)

	delete(mapInit, "t")
	fmt.Println(mapInit)

	if val, exists := mapInit["a"]; exists {
		fmt.Println("Yes \"a\"", val)
	}

	if _, exists := mapInit["t"]; !exists {
		fmt.Println("No \"t\"")
	}

	mapInit["t"] = "ttt"
	for key, val := range mapInit {
		println(key, ":", val)
	}
}
