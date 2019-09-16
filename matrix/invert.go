package matrix

// Invert transposes a matrix of strings, which is then passed to Echo()
func Invert(records [][]string) (string, error) {
	return Echo(transpose(records))
}

// Credit: https://gist.github.com/tanaikech/5cb41424ff8be0fdf19e78d375b6adb8
func transpose(slice [][]string) [][]string {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]string, xl)
	for i := range result {
		result[i] = make([]string, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}
