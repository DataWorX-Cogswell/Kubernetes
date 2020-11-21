package adder

import "errors"

var ErrInvalidSummand = errors.New("Invalid summand")

func Add(x, y int) (int, error) {
	if x <= 0 || y <= 0 {
		return 0, ErrInvalidSummand
	}
	return x + y, nil
}

var ErrInvalidDiv = errors.New("Cannot divide by zero")

func Divide(x, y int) (int, error) {
	if y == 0 {
		return 0, ErrInvalidDiv
	}

	return x / y, nil
}
