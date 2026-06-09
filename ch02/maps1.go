package main
import (
	"fmt"
)

type Entity1 struct{
	Val1 string
	Val2 string
	Val3 string
}

var b map[int]Entity1

func main() {
	b = make(map[int]Entity1)
	b[1] = Entity1{"G", "P", "U"}
	fmt.Println(b)

	c := map[string]Entity1{
		"One": Entity1{"s", "p", "f"},
		"Two": Entity1{"s", "p", "f"},
	}
	fmt.Println(c)
	
	d := map[int]int{
		1: 1,
		2: 2,
	}
	delete(d, 1)
	fmt.Println(d)

	elem, ok := d[2]
	fmt.Println(elem, ok)
}