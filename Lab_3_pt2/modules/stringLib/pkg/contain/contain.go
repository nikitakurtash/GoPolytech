package contain

import (
	"errors"
	"io/ioutil"
)

func Contains(filePath, substring string) (bool, error) {
	if len(substring) == 0 {
		return false, errors.New("Введите подстроку")
	}
	if len(filePath) == 0 {
		return false, errors.New("Введите путь к файлу")
	}

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return false, err
	}

	sub := []byte(substring)
	end := len(content) - len(sub)
	for i := 0; i <= end; i++ {
		if contains(content[i:i+len(sub)], sub) {
			return true, nil
		}
	}
	return false, nil
}

func contains(a, b []byte) bool {
	for i, j := range b {
		if a[i] != j {
			return false
		}
	}
	return true
}
