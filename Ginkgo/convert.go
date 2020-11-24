package convert

import (
	"errors"
	"strconv"
)

var ErrInvalidString = errors.New("String does not convert to known type")

func Convert(s string) (interface{}, error) {
	u, err := strconv.ParseUint(s, 10, 64)
	if err == nil {
		return u, nil
	}

	i, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return i, nil
	}

	f, err := strconv.ParseFloat(s, 64)
	if err == nil {
		return f, nil
	}

	b, err := strconv.ParseBool(s)
	if err == nil {
		return b, nil
	}

	return 0, ErrInvalidString
}
