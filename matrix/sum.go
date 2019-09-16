package matrix

import (
	"errors"
	"strconv"

	"github.com/JohnCGriffin/overflow"
)

// Sum ...
func Sum(records [][]string) (string, error) {
	var sum int
	for _, row := range records {
		rowSum, err := sumArray(row)
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

func sumArray(arr []string) (int, error) {
	var sum int
	for _, x := range arr {
		parsed, err := strconv.Atoi(x)
		if err != nil {
			return 0, err
		}
		var ok bool
		sum, ok = overflow.Add(sum, parsed)
		if !ok {
			return 0, errors.New("Overflow detected")
		}
	}
	return sum, nil
}
