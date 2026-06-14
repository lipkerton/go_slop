package main
import "fmt"

type Entity struct {
    Name        string
    Surname     string
}

func NewEntity(name string, surname string) *Entity {
    return &Entity{Name: name, Surname: surname}
}

func (e *Entity) getFirstName() string {
    return string(e.Name[0])
}

func (e *Entity) getFirstSurname() string {
    return string(e.Surname[0])
}

type EntityName interface {
   getFirstName() string
}

func testInterStruct(e EntityName) string {
    return e.getFirstName()
}

func main() {
    var e *Entity = NewEntity("Peter", "Petukhov")
    fmt.Println(testInterStruct(e))
}
