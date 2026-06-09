package main
import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Printf("EqualFold: %v\n", strings.EqualFold("Cool", "COol"))
	
	fmt.Printf("Index: %v\n", strings.Index("Cool", "o"))
	
	fmt.Printf("Prefix: %v\n", strings.HasPrefix("Cool", "Co"))

	t := string.Fields("String with spaces!")
	fmt.Println("Fields: %v\n", t)

	fmt.Printf("%s\n", strings.Split("abcd efg", ""))
	
	fmt.Printf("%s\n", strings.Replace("abcd efg"))
}