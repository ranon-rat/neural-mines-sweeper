package core

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// first output and then the inpu
// output,input
func LoadData(filename string, outputLength int, maxValueInput, maxValueOutput float32) (input [][]float32, expected [][]float32) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	records, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}
	for _, row := range records {
		inputRow := []float32{}
		outputRow := []float32{}
		for i, v := range row {
			s, _ := strconv.ParseFloat(v, 32)

			if i < (outputLength) {
				outputRow = append(outputRow, float32(s)/maxValueOutput)

				continue
			}
			inputRow = append(inputRow, float32(s)/maxValueInput)

		}

		input = append(input, inputRow)
		expected = append(expected, outputRow)
	}

	return
}

func CreateData(filename string, input, output [][]float32, add bool) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil || !add {
		f, _ = os.Create(filename)
	}
	csv := csv.NewWriter(f)

	defer f.Close()
	for i := range input {
		row := []string{}
		for _, k := range output[i] {
			row = append(row, fmt.Sprintf("%d", int(k)))

		}
		for _, k := range input[i] {
			row = append(row, fmt.Sprintf("%d", int(k)))

		}
		csv.Write(row)

	}
	csv.Flush()

}
