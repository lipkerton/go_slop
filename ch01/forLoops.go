package main
import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Print(i * i, " ")
	}
	fmt.Println()
	for i, j := 0, 0; i < 10; i, j = i + 1, j + 1 {
		fmt.Print(i * j, " ")
	}
	fmt.Println()
	var i int
	for ; ; i++ {
		if i >= 10 {
			break
		}
		fmt.Print(i * i, " ")
	}
	fmt.Println()
	j := 0
	for ok := true; ok; ok = (j != 10) {
		fmt.Print(j * j, " ")
		j++
	}
	fmt.Println()
	k := 0
	for {
		if k == 10 {
			break
		}
		fmt.Print(k * k, " ")
		k++
	}
	fmt.Println()
	aSlice := []int{-1, 2, -1, 2, -2}
	for i, v := range aSlice {
		fmt.Println("index:", i, "value:", v)
	}
}