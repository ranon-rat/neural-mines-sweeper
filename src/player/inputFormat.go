package player

import "github.com/ranon-rat/neural-mines-sweeper/src/core"

type CalfAndPos struct {
	Pos  core.XY
	Calf float32
}

func getLAndR(visibleBoard [][]int, y, x, row int, out []float32, scale float32) []float32 {
	out[row*3+1] = float32(visibleBoard[y][x]) / (scale)
	if x != 0 {
		out[row*3] = float32(visibleBoard[y][x-1]) / scale
	}
	if x != len(visibleBoard[0])-1 {
		out[row*3+2] = float32(visibleBoard[y][x+1]) / scale
	}
	return out

}
func GetInput(visibleBoard [][]int, y, x int, scale float32) (out []float32) {
	out = []float32{-1 / scale, -1 / scale, -1 / scale,
		-1 / scale, -1 / scale, -1 / scale,
		-1 / scale, -1 / scale, -1 / scale}
	out = getLAndR(visibleBoard, y, x, 1, out, scale)
	if y != 0 {
		out = getLAndR(visibleBoard, y-1, x, 0, out, scale)
	}
	if y != len(visibleBoard)-1 {
		out = getLAndR(visibleBoard, y+1, x, 2, out, scale)
	}

	return

}

func GetBestPos(a []CalfAndPos) (index int) {
	for i := 0; i < len(a); i++ {
		if a[index].Calf > a[i].Calf {
			continue
		}
		index = i
	}
	return

}
