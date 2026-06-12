package main
import "fmt"


type aStructure struct {
	field1 complex128
	field2 int
}

func processPointer(x *float64) {
	*x = *x * *x
}

func returnPointer(x float64) *float64 {
	temp := x * 2
	return &temp
}

func bothPointers(x *float64) *float64 {
	temp := 2 * *x
	return &temp
}

func main() {
	var a float64 = 5
	processPointer(&a)
	temp1 := returnPointer(a)
	temp2 := bothPointers(&a)
	fmt.Println(a, *temp1, *temp2)

	var f float64 = 12.123
	fmt.Println("Memory address:", &f)

	fp := &f
	fmt.Println("Memory address:", fp)
	fmt.Println("Value of f:", *fp)
	processPointer(fp)
	fmt.Printf("Value of f: %.2f\n", f)

	var k *aStructure
	fmt.Println(k)

	if k == nil {
		k = new(aStructure)
	}
	fmt.Println(k)
	fmt.Println(*k)
}