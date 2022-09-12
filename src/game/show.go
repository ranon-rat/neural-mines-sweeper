package game

import (
	"fmt"
)

func ShowBoard(board [][]int) {
	fmt.Print(" ")
	for i := 0; i < len(board[0]); i++ {
		fmt.Print("_ ")
	}
	fmt.Println()
	for y := 0; y < len(board); y++ {
		fmt.Print(" | ")
		for x := 0; x < len(board[y]); x++ {

			if _, e := Characters[board[y][x]]; e {
				fmt.Print(Characters[board[y][x]])
				continue
			}
			fmt.Printf("%d ", board[y][x]-1)
		}
		fmt.Println()
	}
	fmt.Println()
}
