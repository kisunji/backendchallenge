package matrix

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
