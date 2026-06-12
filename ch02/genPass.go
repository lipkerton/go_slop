package main
import (
	"fmt"
	"math/rand"
)

const MIN, MAX int = 0, 94

func random(min, max int) int {
	return rand.Intn(max - min) + min
}

func genPass(len int64) string {
	temp := ""
	startChar := "!"
	var i int64 = 1
	for {
		myRand := random(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == len {
			break
		}
		i++
	}
	return temp
}

func main() {
	fmt.Println(genPass(10))
}