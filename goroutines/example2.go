package main
import (
	"fmt"
	"sync"
	"time"
)

func myFunc(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("Выполняю работу %d\n", i)
		time.Sleep(100 * time.Microsecond)		
	}
}

func main() {
	const workers int = 5
	var wg sync.WaitGroup
	for i := 1; i <= workers; i++ {
		wg.Add(1)
		go myFunc(&wg)
	}
	wg.Wait()
}