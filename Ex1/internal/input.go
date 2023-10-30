package internal

import (
	"errors"
	"fmt"
	"os"
)

func GetNumber(message string) (int, error) {
	var result int
	fmt.Print(message)
	_, err := fmt.Fscan(os.Stdin, &result)

	if err != nil {
		return 0, errors.New("Некорректное число. Пожалуйста, введите числовое значение.")
	}
	return result, nil
}
