package main
import "fmt"

type Entity1 struct {
	Val1 string
	Val2 string
	Val3 string
}

func main() {
	var a Entity1
	a.Val1 = "asdfsdf"
	a.Val2 = "dsfwdf"
	a.Val3 = "dfsdfsd"
	fmt.Println(a)

	var p *Entity1
	p = &a
	fmt.Println(p.Val1)
	fmt.Println(p.Val2)
	fmt.Println(p.Val3)
}