package main

import (
	"fmt"
	"strings"
)

func main() {
	// FEN := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	FEN := "rnbqkbnr/ppp2ppp/3p4/4p3/3PP3/5N2/PPP2PPP/RNBQKB1R b KQkq d3 0 3"
	board := strings.Split(FEN, "")
	fmt.Println(board)

	for i := 0; i < len(board); i++ {
		if board[i] == "/" {
			fmt.Print("\n")
		} else if board[i] == "1" || board[i] == "2" || board[i] == "3" || board[i] == "4" || board[i] == "5" || board[i] == "6" || board[i] == "7" || board[i] == "8" {
			fmt.Print(strings.Repeat(" ", int(board[i][0])-48))
		} else {
			fmt.Print(board[i])
		}
	}
}
