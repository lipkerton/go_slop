package main
import (
	"os"
	"fmt"
	"strconv"
)

func main() {
	argument := os.Args
	if len(argument) == 1 {
		fmt.Println("введите аргумент")
		return
	}
	number, err := strconv.Atoi(argument[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	input := strconv.Itoa(number)
	fmt.Println("strconv.Itoa()", input, "of type string")
	input = strconv.FormatInt(int64(number), 10)
	fmt.Println("strconv.FormatInt()", input, "of type string")
	input = string(number)
	fmt.Println("string()", input, "of type string")

	var a int64 = 123
	fmt.Println(strconv.FormatInt(a, 2))
	fmt.Println(strconv.FormatInt(a, 10))
}