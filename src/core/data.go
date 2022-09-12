package core

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// first output and then the input
// the output that it returns you , it comes in this format
// for example
// 5->[0,0,0,0,0,1]
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
	// I shuffle the data, for the training process
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(records), func(i, j int) { records[i], records[j] = records[j], records[i] })

	for _, row := range records {
		inputRow := []float32{}
		// for this project i need to do this
		out := make([]float32, int(maxValueOutput+1))
		for i, v := range row {
			s, _ := strconv.ParseFloat(v, 32)

			if i < outputLength {
				out[int(s)] = 1 // because its a simple classifier Im making this
				continue
			}

			inputRow = append(inputRow, float32(s)/maxValueInput)

		}

		input = append(input, inputRow)
		expected = append(expected, out)
	}

	return
}

func CreateData(filename string, input, output [][]float32, add bool, scale float32) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644) // Sometimes I want to add new stuff , this is why this exist
	if err != nil || !add {
		f, _ = os.Create(filename)
	}
	csv := csv.NewWriter(f)

	defer f.Close()
	for i := range input {

		row := []string{}
		row = append(row, fmt.Sprintf("%d", GetBiggerIndex(output[i])))
		for _, k := range input[i] {
			row = append(row, fmt.Sprintf("%d", int(k*scale)))

		}
		csv.Write(row)

	}
	csv.Flush()

}
