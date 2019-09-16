package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/kisunji/backendchallenge/matrix"
)

func main() {
	http.HandleFunc("/echo", handleEcho())
	http.HandleFunc("/invert", handleInvert())
	http.HandleFunc("/flatten", handleFlatten())
	http.HandleFunc("/sum", handleSum())
	http.HandleFunc("/multiply", handleMultiply())
	http.ListenAndServe(":8080", nil)
}

func handleEcho() http.HandlerFunc {
	return handleCSV(matrix.Echo)
}

func handleInvert() http.HandlerFunc {
	return handleCSV(matrix.Invert)
}

func handleFlatten() http.HandlerFunc {
	return handleCSV(matrix.Flatten)
}

func handleSum() http.HandlerFunc {
	return handleCSV(matrix.Sum)
}

func handleMultiply() http.HandlerFunc {
	return handleCSV(matrix.Multiply)
}

func handleCSV(op matrix.Operation) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		records, err := extractCSV(r)
		if err != nil {
			io.WriteString(w, fmt.Sprintf("CSV parsing error: %s", err.Error()))
			return
		}
		response, err := op(records)
		if err != nil {
			io.WriteString(w, fmt.Sprintf("error %s", err.Error()))
			return
		}
		io.WriteString(w, response)
	}
}

func extractCSV(r *http.Request) ([][]string, error) {
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	if !strings.HasSuffix(fileHeader.Filename, ".csv") {
		return nil, errors.New("File is not .csv")
	}

	return csv.NewReader(file).ReadAll()
}
