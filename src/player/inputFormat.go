package player

import (
	"github.com/ranon-rat/neural-mines-sweeper/src/core"
	"github.com/ranon-rat/neural-mines-sweeper/src/game"
)

type CalfAndPos struct {
	Pos  core.XY
	Calf []float32
}

func getLAndR(visibleBoard [][]int, y, x, row int, out []float32, scale float32) []float32 {
	if x-1 > 0 {
		out[row*5] = float32(visibleBoard[y][x-2]) / scale

	}
	if x > 0 {
		out[row*5+1] = float32(visibleBoard[y][x-1]) / scale
	}
	out[row*5+2] = float32(visibleBoard[y][x]) / (scale)

	if x < len(visibleBoard[0])-1 {
		out[row*5+3] = float32(visibleBoard[y][x+1]) / scale
	}
	if x+1 < len(visibleBoard[0])-1 {
		out[row*5+4] = float32(visibleBoard[y][x+2]) / scale
	}
	return out

}

// its sjust a way of getting the output
// i know that this isnt beautiful and optimize
// way of doing this
// but im to lazy for implementing something else
func GetInput(visibleBoard [][]int, y, x int, scale float32) (out []float32) {
	out = []float32{
		game.UndiscoveredCell / scale, game.UndiscoveredCell / scale, game.UndiscoveredCell / scale, game.UndiscoveredCell / scale, game.UndiscoveredCell / scale,
		game.UndiscoveredCell / scale, game.UndiscoveredCell / scale, game.UndiscoveredCell / scale, game.UndiscoveredCell / scale, game.UndiscoveredCell / scale,
		game.UndiscoveredCell / scale, game.UndiscoveredCell / scale, game.UndiscoveredCell / scale, game.UndiscoveredCell / scale, game.UndiscoveredCell / scale,
		game.UndiscoveredCell / scale, game.UndiscoveredCell / scale, game.UndiscoveredCell / scale, game.UndiscoveredCell / scale, game.UndiscoveredCell / scale,
		game.UndiscoveredCell / scale, game.UndiscoveredCell / scale, game.UndiscoveredCell / scale, game.UndiscoveredCell / scale, game.UndiscoveredCell / scale,
	}
	if y-1 > 0 {
		out = getLAndR(visibleBoard, y-2, x, 0, out, scale)
	}
	if y > 0 {
		out = getLAndR(visibleBoard, y-1, x, 1, out, scale)
	}

	out = getLAndR(visibleBoard, y, x, 2, out, scale)

	if y < len(visibleBoard)-1 {
		out = getLAndR(visibleBoard, y+1, x, 3, out, scale)
	}
	if y+1 < len(visibleBoard)-1 {
		out = getLAndR(visibleBoard, y+2, x, 4, out, scale)

	}
	return

}

// it just get me the bigger value
func GetBestPos(a []CalfAndPos) (index int) {
	for i := 0; i < len(a); i++ {
		if a[index].Calf[0] > a[i].Calf[0] {
			continue
		}
		index = i
	}
	return

}
