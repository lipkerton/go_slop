package main
import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		fmt.Println("Горутина отправляет 42!")
		ch <- 42
		fmt.Println("Горутина продолжила после отправки!")
	}()

	fmt.Println("Main ждет данные...")
	value := <-ch
	fmt.Println("Горутина прислала данные:", value)
}