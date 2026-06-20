package main

func f() (x int) {
	defer func() {
		x = 2
	}()

	return 1
}
