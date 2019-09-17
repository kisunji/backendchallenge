package matrix

import (
	"errors"
	"strconv"

	"github.com/JohnCGriffin/overflow"
)

// Sum returns the sum of all the parsed integers in records
// Returns error if it cannot parse string->int OR if there is int overflow
func Sum(records [][]string) (string, error) {
	intMatrix, err := AtoiMatrix(records)
	if err != nil {
		return "", err
	}
	var sum int
	for _, row := range intMatrix {
		rowSum, err := sumSlice(row)
		if err != nil {
			return "", err
		}
		var ok bool
		sum, ok = overflow.Add(sum, rowSum)
		if !ok {
			return "", errors.New("Overflow detected")
		}
	}
	return strconv.Itoa(sum), nil
}

func sumSlice(arr []int) (int, error) {
	var sum int
	for _, x := range arr {
		var ok bool
		sum, ok = overflow.Add(sum, x)
		if !ok {
			return 0, errors.New("Overflow detected")
		}
	}
	return sum, nil
}
