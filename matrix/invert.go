package matrix

// Invert transposes a matrix of strings, then passes result to to Echo()
func Invert(records [][]string) (string, error) {
	return Echo(transpose(records))
}

// Credit: https://gist.github.com/tanaikech/5cb41424ff8be0fdf19e78d375b6adb8
func transpose(records [][]string) [][]string {
	if len(records) < 1 {
		return nil
	}
	xl := len(records[0])
	yl := len(records)
	result := make([][]string, xl)
	for i := range result {
		result[i] = make([]string, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = records[j][i]
		}
	}
	return result
}
