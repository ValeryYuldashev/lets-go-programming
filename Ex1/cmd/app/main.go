package main

import (
	"Ex1/internal"
	"fmt"
	"os"
)

func main() {
	var operation string

	a, err := internal.GetNumber("Введите первое число: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("Выберите операцию (+, -, *, /): ")
	fmt.Fscan(os.Stdin, &operation)

	b, err := internal.GetNumber("Введите второе число: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	result, err := internal.Calculate(a, b, operation)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(fmt.Sprintf("Результат: %d %s %d = %d", a, operation, b, result))
}
