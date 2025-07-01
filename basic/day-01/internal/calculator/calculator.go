package calculator

import "errors"

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) (int, error) {
	if condition := b == 0; condition {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func AddAndSubtract(a, b int) (add, subtract int) {
	add = Add(a, b)
	subtract = Subtract(a, b)
	return
}
