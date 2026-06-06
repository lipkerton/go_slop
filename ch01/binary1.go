package main
import (
	"fmt"
)


func sumTwoBinary(a string, b string) []byte {
	len_a := len(a) - 1
	len_b := len(b) - 1
	carry := 0
	result := []byte{}
	
	for ; len_a >= 0 || len_b >= 0 || carry > 0; len_a, len_b = len_a - 1, len_b - 1 {
		var digit_a, digit_b byte = 0, 0
		if len_a >= 0 {
			digit_a = a[len_a]
		}
		if len_b >= 0 {
			digit_b = b[len_b]
		}
		sum := int(digit_a - '0') + int(digit_b - '0') + carry
		switch sum {
		case 0:
			result = append(result, 0)
		case 1:
			result = append(result, 1)
		case 2:
			result = append(result, 0)
			carry += 1
		case 3:
			result = append(result, 1)
			carry += 1
		}
		carry = sum / 2
	}
	return result
}


func main() {
	var a, b string
	fmt.Scan(&a, &b)
	fmt.Println(sumTwoBinary(a, b))
}