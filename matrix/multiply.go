package matrix

import (
	"errors"
	"strconv"

	"github.com/JohnCGriffin/overflow"
)

// Multiply returns the product of all the parsed integers in records
// Returns error if it cannot parse string->int OR if there is int overflow
func Multiply(records [][]string) (string, error) {
	if len(records) == 0 {
		return "0", nil
	}
	intMatrix, err := AtoiMatrix(records)
	if err != nil {
		return "", err
	}
	matrixProduct := 1
	for _, row := range intMatrix {
		rowProduct, err := productSlice(row)
		if err != nil {
			return "", err
		}
		var ok bool
		matrixProduct, ok = overflow.Mul(matrixProduct, rowProduct)
		if !ok {
			return "", errors.New("Overflow detected")
		}
	}
	return strconv.Itoa(matrixProduct), nil
}

func productSlice(arr []int) (int, error) {
	if len(arr) == 0 {
		return 0, nil
	}
	product := 1
	for _, x := range arr {
		var ok bool
		product, ok = overflow.Mul(product, x)
		if !ok {
			return 0, errors.New("Overflow detected")
		}
	}
	return product, nil
}
