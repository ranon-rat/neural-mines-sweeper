package player

import (
	"github.com/ranon-rat/neural-mines-sweeper/src/core"
)

// o my fucking god this is fucking trash aaaaa
func FindAvaibleCells(visibleBoard [][]int) (xyAvaible []core.XY) {

	for i, v := range visibleBoard {
		for j := range v {
			input := GetInput(visibleBoard, i, j, 1)

			if input[4] == -1 {
				var s float32 = 0
				for _, q := range input {
					if q != -1 {
						continue
					}
					s += (q)
				}
				if s > -7 {

					xyAvaible = append(xyAvaible, core.XY{X: j, Y: i})
				}
			}
		}
	}

	return
}
