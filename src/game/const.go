package game

const (
	bomb             = 9
	UndiscoveredCell = -1
	nothing          = 0
)

var (
	Characters = map[int]string{bomb: "* ", UndiscoveredCell: "# ", nothing: "  "}
)
