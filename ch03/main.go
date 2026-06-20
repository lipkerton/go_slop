package main

import "fmt"

func main() {
	var m map[string]int = map[string]int{}
	myMapp1(m, "one", 1)
	fmt.Println(m)

	var n myStruct1
	var b []myStruct1

	n = myStruct1{a: "Привет!", b: 3}
	b = []myStruct1{
		myStruct1{a: "Привет!", b: 1},
		myStruct1{a: "Пока!", b: 2},
	}
	fmt.Println(n, b)

	fmt.Println(f())
}
