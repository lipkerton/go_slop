package main
import (
	"fmt"
)

func main() {
	b := make(chan int, 3)
	b <- 1
	b <- 2
	b <- 3
	close(b)  // закрытие канала

	for v := range b {
		fmt.Println(v)
	}

	v, ok := <-b
	fmt.Println(v, ok)
}