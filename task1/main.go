package main

import (
	"fmt"
)

func main() {
	funcWithPanic()
}

func funcWithPanic() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Printf("Поймана паника: %s\n", v)
		}
	}()

	var a int = 0
	_ = 1 / a
}
