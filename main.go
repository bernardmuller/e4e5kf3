package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// FEN := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	FEN := "rnbqkbnr/ppp2ppp/3p4/4p3/3PP3/5N2/PPP2PPP/RNBQKB1R b KQkq d3 0 3"
	// FEN := "r4rk1/1pp1nppp/p1nqp3/1B1pNb2/3P1P2/2P5/PP1N1PPP/R2Q1RK1 w - - 0 11"
	FEN_layout := strings.Split(strings.Split(FEN, " ")[0], "")

	row := 0
	col := 0

	current_board := make([][]string, 8)
	for i := range current_board {
		current_board[i] = make([]string, 8)
	}

	for i := 0; i < len(FEN_layout); i++ {
		if FEN_layout[i] == "/" {
			row++
			col = 0
		} else {
			if FEN_layout[i] == "1" || FEN_layout[i] == "2" || FEN_layout[i] == "3" || FEN_layout[i] == "4" ||
				FEN_layout[i] == "5" || FEN_layout[i] == "6" || FEN_layout[i] == "7" || FEN_layout[i] == "8" {
				number, _ := strconv.Atoi(string(FEN_layout[i]))

				for x := 0; x < number; x++ {
					current_board[row][col] = " "
					col++
				}
			} else {
				current_board[row][col] = string(FEN_layout[i])
				col++
			}
		}
	}

	for i := 0; i < 8; i++ {
		fmt.Print(8 - i)
		fmt.Print(" ")
		for j := 0; j < 8; j++ {
			fmt.Print(current_board[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Print("  abcdefgh\n")

	// chessboard := [8][8]string{
	// 	{"a8", "b8", "c8", "d8", "e8", "f8", "g8", "h8"},
	// 	{"a7", "b7", "c7", "d7", "e7", "f7", "g7", "h7"},
	// 	{"a6", "b6", "c6", "d6", "e6", "f6", "g6", "h6"},
	// 	{"a5", "b5", "c5", "d5", "e5", "f5", "g5", "h5"},
	// 	{"a4", "b4", "c4", "d4", "e4", "f4", "g4", "h4"},
	// 	{"a3", "b3", "c3", "d3", "e3", "f3", "g3", "h3"},
	// 	{"a2", "b2", "c2", "d2", "e2", "f2", "g2", "h2"},
	// 	{"a1", "b1", "c1", "d1", "e1", "f1", "g1", "h1"},
	// }
}
