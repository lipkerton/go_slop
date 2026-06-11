package main
import (
	"fmt"
)

func main() {
	var target int = 2
	var d []int = []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("Удаляемый элемент %d по индексу %d\n", d[target], target)
	fmt.Println()
	fmt.Println("===СПОСОБ №1===")
	s_1, s_2 := d[:target], d[target + 1:]
	result := append(s_1, s_2...)
	fmt.Printf("Результат удаления: %v\n", result)
	fmt.Println()
	fmt.Println("===СПОСОБ №2===")
	d[2] = d[len(d) - 1]
	result = d[:len(d) - 2]
	fmt.Printf("Результат удаления: %v\n", result)
}