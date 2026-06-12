package main
import (
	"fmt"
	"os"
	"path"
	"strconv"
	"math/rand"
)

type Entry struct {
	Name string
	Surname string
	Tel string
}

var data []Entry

func getString(len int) string {
	b := make([]byte, len, len)
	for i := 0; i < len; i++ {
		b = append(b, byte(rand.Intn(90 - 65) + 65))
	}
	return string(b)
}

func random(min, max int) int {
	return rand.Intn(max - min) + min
}

func search(key string) *Entry {
	for index, value := range data {
		if value.Surname == key {
			return &data[index]
		}
	}
	return nil
}

func list() {
	for _, v := range data {
		fmt.Println(v)
	}
}

func populate(n int) []Entry {
	s := make([]Entry, n, n)
	for i := 0; i < n; i++ {
		name := getString(4)
		surname := getString(5)
		b := strconv.Itoa(random(100, 199))
		s[i] = Entry{name, surname, b}
	}
	return s
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		exe := path.Base(arguments[0])
		fmt.Println("Usage: %s search|list <arguments>\n", exe)
		return
	}

	data = populate(10)
	switch arguments[1] {
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Usage: search Surname")
			return
		}
		result := search(arguments[2])
		if result == nil {
			fmt.Println("Entry not found:", arguments[2])
			return
		}
		fmt.Println(*result)
	case "list":
		list()
	default:
		fmt.Println("Not valid option!")
	}
}