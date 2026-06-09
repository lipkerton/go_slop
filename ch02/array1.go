package main
import "fmt"

func main() {
	var a [10]int
	a = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(a)

	a[0] = 100
	fmt.Println(a)

	var b []int
	b = a[1 : 3]
	fmt.Println(b)

	c := []struct{
		i string
		j string
	}{
		{"dsf", "fdsf"},
		{"sdfsdf", "fsdfsf"},
	}
	fmt.Println(c)
}