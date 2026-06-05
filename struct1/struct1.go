package main
import (
	"fmt"
	"struct1/models"
)

func main() {
	var a models.Entity
	var b *models.Entity
	a = models.Entity{Name: "Slavoy", Surname: "Zijek"}
	b = &a
	fmt.Println(a)
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", b)
}