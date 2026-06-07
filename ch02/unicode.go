package main
import (
	"fmt"
	"unicode"
)

func main() {
	var a []rune = []rune("\x99\x00ab\x50\x00\x23\x50\x29\x9c")
	for i := 0; i < len(a); i++ {
		if unicode.IsPrint(a[i]) {
			fmt.Printf("%c\n", a[i])
		} else {
			fmt.Println("Not printable!")
		}
	}
}