package main
import (
	"os"
	"fmt"
	"strconv"
)

func main() {
	arguments := os.Args
	var total, nInts, nFloats int
	invalid := make([]string, 0, 3)
	for _, k := range arguments[1:] {
		_, err := strconv.Atoi(k)
		if err == nil {
			nInts++
			total++
			continue
		}
		_, err = strconv.ParseFloat(k, 64)
		if err == nil {
			nFloats++
			total++
			continue
		}
		invalid = append(invalid, k)
	}
	fmt.Println(invalid)
}