package main
import (
	"fmt"
	"errors"
)

func main() {
	var a, b error
	a = errors.New("ошибка")
	b = fmt.Errorf("ошибка")
	fmt.Println(a, b)
}