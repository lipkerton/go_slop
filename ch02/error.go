package main
import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func check(a, b int) error {
	if a == b && b == 0 {
		return error.New("this is custom error message")
	}
	return nil
}