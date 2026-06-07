package main
import (
	"fmt"
)

func main() {
	aString := "Hello world! &"
	fmt.Println("First character", string(aString[0]))

	r := 'Ц'
	fmt.Println("As an int32 value:", r)
	fmt.Printf("As a string: %s and as a character: %c\n", r, r)

	for _, v := range aString {
		fmt.Printf("%c", v)
	}
	fmt.Println()
}