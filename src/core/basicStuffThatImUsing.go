package core

type XY struct {
	X int
	Y int
}
type UniquePosition map[XY]bool

func (u UniquePosition) Add(val XY) {
	if u[val] {
		return
	}
	u[val] = true

}

func GetBiggerIndex(input []float64) (index int) {
	for i := 0; i < len(input)-1; i++ {
		if input[index] > input[i] {
			continue

		}
		index = i
	}
	return
}
