package calculations

import "fmt"

func Calculate(a int64, log bool) int64 {
	if log {
		fmt.Printf("Start calculations...\nCalculate %d!\n", a)
	}
	var factorial int64 = 1
	for i := int64(1); i <= a; i++ {
		factorial *= i
	}
	if log {
		fmt.Println("Calculations complete!")
	}
	return factorial
}
