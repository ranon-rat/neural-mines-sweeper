package player

import "github.com/ranon-rat/neural-mines-sweeper/src/core"

type CalfAndPos struct {
	Pos  core.XY
	Calf float64
}

func getLAndR(visibleBoard [][]int, y, x, row int, input []float64) (out []float64) {
	out = input
	out[row*3+1] = float64(visibleBoard[y][x]) / 9
	if x != 0 {
		out[row*3] = float64(visibleBoard[y][x-1]) / 9
	}
	if x != len(visibleBoard[0])-1 {
		out[row*3+2] = float64(visibleBoard[y][x+1]) / 9
	}
	return

}
func getInput(visibleBoard [][]int, y, x int) (out []float64) {
	out = []float64{-1 / 9, -1 / 9, -1 / 9,
		-1 / 9, -1 / 9, -1 / 9,
		-1 / 9, -1 / 9, -1 / 9}
	out = getLAndR(visibleBoard, y, x, 1, out)
	if y != 0 {
		out = getLAndR(visibleBoard, y-1, x, 0, out)
	}
	if y != len(visibleBoard)-1 {
		out = getLAndR(visibleBoard, y+1, x, 2, out)
	}

	return

}

func GetBestPos(a []CalfAndPos) (index int) {
	for i := 0; i < len(a)-1; i++ {
		if a[index].Calf > a[i].Calf {
			continue
		}
		index = i
	}
	return

}
