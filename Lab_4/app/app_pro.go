//go:build pro
// +build pro

package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 1
	b := 2
	fmt.Println("В Про версии я могу сложить два числа: " + strconv.Itoa(add(a, b)))
}

func add(a, b int) int {
	c := a + b
	return c
}
