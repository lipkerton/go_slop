package main
import "fmt"


func eval() func() int {
	counter := 0
	return func() int {
		counter++
		return counter
	}
}

func main() {
	c1 := eval()
	c2 := eval()

	fmt.Println(c1())
	fmt.Println(c2())
	fmt.Println(c1())
}