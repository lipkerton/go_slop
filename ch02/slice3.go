package main
import (
	"fmt"
)

func main() {
	a := make([]int, 4)
	fmt.Println("L:", len(a), "C:", cap(a))

	a = append(a, 0)
	fmt.Println("L:", len(a), "C:", cap(a))

	var b, d []int
	b = []int{1, 2, 3, 4}
	d = []int{1, 2, 3}
	b = append(b, d...)
	fmt.Println(b)
}