package matrix

import "strconv"

// Operation is a function that is applied to a matrix of strings to produce an output string and/or error
type Operation func(matrix [][]string) (string, error)

// IsSquareMatrix can be used to validate that a matrix is N*N
func IsSquareMatrix(records [][]string) bool {
	for y := 0; y < len(records); y++ {
		colLength := len(records)
		for x := 0; x < len(records[y]); x++ {
			rowLength := len(records[y])
			if colLength != rowLength {
				return false
			}
		}
	}
	return true
}

// AtoiMatrix applies AtoiSlice to each slice in the input
func AtoiMatrix(records [][]string) ([][]int, error) {
	intMatrix := make([][]int, len(records))
	for i, row := range records {
		intSlice, err := AtoiSlice(row)
		if err != nil {
			return nil, err
		}
		intMatrix[i] = intSlice
	}
	return intMatrix, nil
}

// AtoiSlice applies strconv.Atoi to each element in in the input
func AtoiSlice(slice []string) ([]int, error) {
	intSlice := make([]int, len(slice))
	for i, x := range slice {
		val, err := strconv.Atoi(x)
		if err != nil {
			return nil, err
		}
		intSlice[i] = val
	}
	return intSlice, nil
}
