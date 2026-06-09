package main
import (
	"fmt"
)

func addOne(x int, y int) int {
	return x + y
}

func swap(x string, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(addOne(1, 1))
	fmt.Println(swap("Hello world!", "Goodbye world!"))
	fmt.Println(split(12312))

	var a int = 10
	b := float32(a)
	fmt.Println(b)
}