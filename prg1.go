package main
import (
	"fmt"
	"time"
)

func subFunc() {
	for i := 1; i <= 3; i++ {
		fmt.Println("Привет из дочернего потока!", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go subFunc()

	for i := 1; i <= 3; i++ {
		fmt.Println("Привет из основного потока!", i)
		time.Sleep(100 * time.Millisecond)
	}

	time.Sleep(200 * time.Millisecond)
	fmt.Println("Программа завершена!")
}