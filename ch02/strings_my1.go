package main
import (
	"fmt"
)

func main() {
	var a string = "A string"
	var b []byte = []byte(a)
	fmt.Println(a)
	fmt.Println(string(b))
}