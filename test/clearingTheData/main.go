package main

import "github.com/ranon-rat/neural-mines-sweeper/src/core"

func main() {
	newDataset, newExpected := [][]float32{}, [][]float32{}
	dataset, expected := core.LoadData("../../data/minessweeper.csv", 1, 1, 1)
	bomb := 0
	for j, v := range expected {

		if v[0] == 1 && bomb > 0 {
			newDataset = append(newDataset, dataset[j])
			newExpected = append(newExpected, v)
			bomb--
		} else if v[0] == 0 && bomb == 0 {
			newDataset = append(newDataset, dataset[j])
			newExpected = append(newExpected, v)
			bomb++
		}

	}
	core.CreateData("../../data/minessweeper.csv", newDataset, newExpected, false, 1)

}
