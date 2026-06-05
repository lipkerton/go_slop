package main
import (
	"fmt"
)

type Entry struct {
	Name string
	Surname string
	Tel string
}

var data []Entry{}

func search(key string) *Entry {
	for index, value := range data {
		if value.Surname == key {
			return &data[i]
		}
	}
	return nil
}

func main() {
	var a []Entry
	a = []Entry{Entry{Name: "Peter", Surname: "Tuturin", Tel: "1231345"},
		Entry{Name: "Olesya", Surname: "Tuturina", Tel: "131235"}}
	fmt.Println(a)
}