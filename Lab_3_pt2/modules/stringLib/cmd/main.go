package main

import (
	"github.com/nikitakurtash/stringLib/pkg/contain"
)

func main() {
	// Пример использования функции
	found, err := contain.Contains("/home/xs/GolandProjects/Lab_3_pt2/modules/stringLib/cmd/file.txt", "polytech")
	if err != nil {
		panic(err) // или обработать ошибку более тонко
	}
	if found {
		println("Substring found!")
	} else {
		println("Substring not found.")
	}
}
