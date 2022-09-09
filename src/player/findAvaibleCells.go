package player

import (
	"github.com/ranon-rat/neural-mines-sweeper/src/core"
)

// o my fucking god this is fucking trash aaaaa
func FindAvaibleCells(xMv, yMv int, visibleBoard [][]int) (xyAvaible []core.XY) {

	for i, v := range visibleBoard {
		for j := range v {
			input := getInput(visibleBoard, i, j)

			if int(input[4]*9) == -1 {
				s := 0
				for _, q := range input {

					s += int(q * 9)
				}
				if s > -8 {

					xyAvaible = append(xyAvaible, core.XY{X: j, Y: i})
				}
			}
		}
	}

	return
}
