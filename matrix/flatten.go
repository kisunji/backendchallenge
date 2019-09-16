package matrix

import (
	"fmt"
	"strings"
)

// Flatten matrix of strings into one line separated by commas
// Works with any N*M matrix
func Flatten(records [][]string) (string, error) {
	var response string
	for _, row := range records {
		response = fmt.Sprintf("%s,%s", response, strings.Join(row, ","))
	}
	return strings.TrimPrefix(response, ","), nil
}
