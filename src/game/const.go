package game

// sorry but i want to use relu
const (
	Bomb             = 10
	UndiscoveredCell = 0
	Nothing          = 1
)

var (
	Characters = map[int]string{Bomb: "* ", UndiscoveredCell: "# ", Nothing: "  "}
)
