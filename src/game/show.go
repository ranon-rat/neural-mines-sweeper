package game

import (
	"fmt"
)

func ShowBoard(board [][]int) {

	fmt.Print("  _")

	for i := 0; i < len(board[0]); i++ {
		fmt.Printf("%d_", i)
	}
	fmt.Println()
	for y := 0; y < len(board); y++ {

		fmt.Printf("%d |", y)

		for x := 0; x < len(board[y]); x++ {

			if board[y][x] == bomb || board[y][x] == UndiscoveredCell || board[y][x] == nothing {
				fmt.Print(Characters[board[y][x]])
				continue
			}
			fmt.Printf("%d ", board[y][x])
		}
		fmt.Println()
	}
	fmt.Println()
}
