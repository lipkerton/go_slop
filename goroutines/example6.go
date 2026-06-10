package main
import (
	"fmt"
	"time"
)

func myFunc(done chan struct{}) {
	fmt.Println("Работаю...")
	time.Sleep(100 * time.Microsecond)
	fmt.Println("Закончил работать!")
	done <- struct{}{}
}

func main() {
	done := make(chan struct{})
	go myFunc(done)
	<-done
	fmt.Println("Main получила сигнал о завершении!")
}
