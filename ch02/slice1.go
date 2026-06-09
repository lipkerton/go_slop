package main
import "fmt"

func main() {
	var a []int = []int{1, 2, 3, 4, 5}
	fmt.Println(a)

	a = a[:0]
	fmt.Println(a)

	a = a[:4]
	fmt.Println(a)
}