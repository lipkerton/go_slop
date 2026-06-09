package main
import "fmt"

func main() {
	defer fmt.Println("мир!")
	fmt.Print("Привет, ")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}