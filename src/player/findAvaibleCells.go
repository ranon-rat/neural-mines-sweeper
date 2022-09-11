package player

import (
	"github.com/ranon-rat/neural-mines-sweeper/src/core"
)

func getArrLAndR(visibleBoard [][]int, y, x, row int, out []int) []int {

	if x > 0 {
		out[row*3] = (visibleBoard[y][x-1])
	}
	out[row*3+1] = (visibleBoard[y][x])

	if x < len(visibleBoard[0])-1 {
		out[row*3+2] = (visibleBoard[y][x+1])
	}

	return out
}
func GetArround(visibleBoard [][]int, y, x int) (out []int) {

	out = []int{
		-1, -1, -1,
		-1, -1, -1,
		-1, -1, -1,
	}

	out = getArrLAndR(visibleBoard, y, x, 1, out)

	if y > 0 {
		out = getArrLAndR(visibleBoard, y-1, x, 0, out)
	}

	if y < len(visibleBoard)-1 {
		out = getArrLAndR(visibleBoard, y+1, x, 2, out)
	}

	return

}

// o my fucking god this is fucking trash aaaaa
func FindAvaibleCells(visibleBoard [][]int) (xyAvaible []core.XY) {
	cells := []core.XY{}

	for i, v := range visibleBoard {
		for j := range v {
			input := GetArround(visibleBoard, i, j)
			if input[4] == -1 {
				cells = append(cells, core.XY{X: j, Y: i})

				s := 0
				for _, q := range input {

					if q != -1 {
						continue
					}
					s += q
				}
				if s > -8 {

					xyAvaible = append(xyAvaible, core.XY{X: j, Y: i})
				}
			}
		}

	}
	if len(xyAvaible) == 0 {
		xyAvaible = cells
	}
	return
}
