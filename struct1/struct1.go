package main
import (
	"fmt"
	"struct1/models"
)

func main() {
	var a models.Entity
	a = models.Entity{Name: "Slavoy", Surname: "Zijek", tel: "dfsdfsdf"}
	fmt.Println(a)
}