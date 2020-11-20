package main

import "fmt"

type person struct {
	name string
	age  int
}

type dictionary struct {
	data map[int]string
}

func newDictionary() *dictionary {
	d := dictionary{}
	d.data = map[int]string{}
	return &d
}

func main() {
	p := person{}
	p.name = "AN"
	p.age = 26

	fmt.Println(p)

	var p1 person
	p1 = person{"test", 20}
	p2 := person{name: "test2", age: 25}
	p3 := person{"test3", 24}
	fmt.Println(p1, p2, p3)

	d := newDictionary()
	d.data[1] = "123"
	d.data[2] = "234"
	fmt.Println(d)
}
