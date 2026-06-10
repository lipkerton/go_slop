package main

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3  // нет блока, но очередь канала заполнена
	ch <- 4  // блок
}