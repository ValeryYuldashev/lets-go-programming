package internal

import "errors"

func Calculate(a int, b int, operation string) (int, error) {
	switch operation {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("Ошибка: деление на ноль невозможно.")
		}
		return a / b, nil
	default:
		return 0, errors.New("Некорректная операция. Пожалуйста, используйте символы +, -, * или /.")
	}
}
