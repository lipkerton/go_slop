package main

import (
	"fmt"
	"math"
)

type Entity struct {
	CarName   string
	CarType   string
	CarNumber int64
}
type Message struct {
	ID      int
	Content string
}

func NewEntity(cname, ctype string, cnum int64) *Entity {
	return &Entity{CarName: cname, CarType: ctype, CarNumber: cnum}
}
func NewMessage(id int, content string) *Message {
	return &Message{ID: id, Content: content}
}

func (e *Entity) getCarNameFirst() string {
	return string(e.CarName[0])
}
func (e *Entity) getCarNumberPowered() int {
	num := float64(e.CarNumber)
	return int(math.Pow(num, 2))
}

func (m *Message) sendMessage(dest string) string {
	return fmt.Sprintf("Your message `%s` was send to %s", m.Content, dest)
}

func main() {
	d := NewEntity("Mersedez", "sedan", 12)
	m := NewMessage(1, "Hello!")
	fmt.Println(d.getCarNameFirst())
	fmt.Println(d.getCarNumberPowered())

	fmt.Println(m.sendMessage("my_friend_228_1488"))
}
