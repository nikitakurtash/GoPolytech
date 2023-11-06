package main

import (
	"fmt"
	"github.com/nikitakurtash/stringLib/pkg/contain"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

func main() {
	var conf Config

	data, err := ioutil.ReadFile("/home/xs/GolandProjects/Lab_3_pt2/cmd/config.yml")
	if err != nil {
		log.Fatalf("Не удалось прочитать файл!")
	}

	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		log.Fatalf("Проверьте формат конфига!")
	}

	for _, fileConfig := range conf.Files {
		contains, err := contain.Contains(fileConfig.Filename, fileConfig.Substring)
		if err != nil {
			log.Printf("Ошибка проверки/чтения файла %s: %v", fileConfig.Filename, err)
			continue
		}
		if contains {
			fmt.Printf("В файле %s найдена подстрока '%s'.\n", fileConfig.Filename, fileConfig.Substring)
		} else {
			fmt.Printf("В файле %s подстрока '%s' не найдена.\n", fileConfig.Filename, fileConfig.Substring)
		}
	}
}
