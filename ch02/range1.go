package main
import "fmt"

func main() {
	var a []int = []int{1, 2, 3, 4, 5}
	for i, v := range a {
		fmt.Println(i, v)
	}
}