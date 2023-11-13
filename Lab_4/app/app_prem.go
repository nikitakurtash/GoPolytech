//go:build prem
// +build prem

package main

import (
	"embed"
	"fmt"
)

//go:embed read.txt
var readedFile string

//go:embed read.txt
var forIDE embed.FS

func main() {
	fmt.Println("В премиум версии я могу прочитать Ваш файл. Вот он: ")
	fmt.Println(readedFile)
}
