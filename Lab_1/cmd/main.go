package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b string
	var aInt, bInt float64
	var o string
	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)
	aInt, e := strconv.ParseFloat(a, 64)
	if e != nil {
		fmt.Println("Некорректное число. Пожалуйста, введите числовое значение.")
		return
	}
	fmt.Print("Выберите операцию (+, -, *, /): ")
	fmt.Scan(&o)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)
	bInt, e1 := strconv.ParseFloat(b, 64)
	if e1 != nil {
		fmt.Println("Некорректное число. Пожалуйста, введите числовое значение.")
		return
	}
	switch o {
	case "+":
		fmt.Printf("Результат: %f + %f = %f", aInt, bInt, aInt+bInt)
	case "-":
		fmt.Printf("Результат: %f - %f = %f", aInt, bInt, aInt-bInt)
	case "*":
		fmt.Printf("Результат: %f * %f = %f", aInt, bInt, aInt*bInt)
	case "/":
		if bInt == 0 {
			fmt.Print("Ошибка: деление на ноль невозможно.")
		} else {
			fmt.Printf("Результат: %f / %f = %.15f", aInt, bInt, aInt/bInt)
		}
	default:
		fmt.Println("Некорректная операция. Пожалуйста, используйте символы +, -, * или /.")
	}
}
