package main
import (
	"fmt"
	"math/rand"
)

func random(min, max int) int {
	return rand.Intn(max - min) + min
}

func main() {
	for i := 0; i < 100; i++ {
		fmt.Printf("%d ", random(1, 100))
	}
	fmt.Println(random(1, 10))
}