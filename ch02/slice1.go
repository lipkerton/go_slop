package main
import "fmt"

func main() {
	var a []int = []int{1, 2, 3, 4, 5}
	fmt.Println(a)

	a = a[:0]
	fmt.Println(a)

	a = a[:4]
	fmt.Println(a)

	var b []int = make([]int, 5, 6)
	b[len(b) - 1] = 5
	fmt.Println(b)

	var c [][]int = [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
	}
	fmt.Println(c)

	c = append(c, []int{7, 8, 9}, []int{10, 11, 12})
	fmt.Println(c)
}