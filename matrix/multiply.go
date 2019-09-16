package matrix

import (
	"errors"
	"strconv"

	"github.com/JohnCGriffin/overflow"
)

// Multiply ...
func Multiply(records [][]string) (string, error) {
	product := 1
	for _, row := range records {
		rowSum, err := productArray(row)
		if err != nil {
			return "", err
		}
		var ok bool
		product, ok = overflow.Mul(product, rowSum)
		if !ok {
			return "", errors.New("Overflow detected")
		}
	}
	return strconv.Itoa(product), nil
}

func productArray(arr []string) (int, error) {
	product := 1
	for _, x := range arr {
		parsed, err := strconv.Atoi(x)
		if err != nil {
			return 0, err
		}
		var ok bool
		product, ok = overflow.Mul(product, parsed)
		if !ok {
			return 0, errors.New("Overflow detected")
		}
	}
	return product, nil
}
