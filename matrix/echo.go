package matrix

import (
	"fmt"
	"strings"
)

// Echo matrix of strings, separating elements within a row with commas and new rows with newline (/n)
// Works with any N*M matrix
func Echo(records [][]string) (string, error) {
	var response string
	for _, row := range records {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	return strings.TrimSuffix(response, "\n"), nil
}
