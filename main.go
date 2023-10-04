package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Piece represents a chess piece.
type Piece string

// display maps pieces from their FEN representations to their ASCII
// representations for a more human readable experience.
var display = map[Piece]string{
	"":  " ",
	"B": "♝",
	"K": "♚",
	"N": "♞",
	"P": "♟",
	"Q": "♛",
	"R": "♜",
	"b": "♗",
	"k": "♔",
	"n": "♘",
	"p": "♙",
	"q": "♕",
	"r": "♖",
}

type File string
type Rank string

var FILES = [8]File{"a", "b", "c", "d", "e", "f", "g", "h"}
var RANKS = [8]Rank{"1", "2", "3", "4", "5", "6", "7", "8"}

func indexOfFile(item File, array [8]File) int {
	for k, v := range array {
		if item == v {
			return k
		}
	}
	return -1
}

func indexOfRank(item Rank, array [8]Rank) int {
	for k, v := range array {
		if item == v {
			return k
		}
	}
	return -1
}

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
				FEN_piece := FEN_layout[i]
				current_board[row][col] = display[Piece(FEN_piece)]
				col++
			}
		}
	}

	for i := 0; i < 8; i++ {
		fmt.Print(8 - i)
		fmt.Print("  ")
		for j := 0; j < 8; j++ {
			fmt.Print(current_board[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Print("   abcdefgh\n")

	var user_input string
	for {
		fmt.Print("Select a square:")
		fmt.Scanf("%s", &user_input)
		fmt.Printf("you selected square %s", user_input)
		fmt.Println("")
		selected_file := strings.Split(user_input, "")[0]
		selected_rank := strings.Split(user_input, "")[1]
		fmt.Printf("%d", indexOfRank(Rank(selected_rank), RANKS))
		fmt.Printf("%d", indexOfFile(File(selected_file), FILES))
		selected_piece := current_board[7-indexOfRank(Rank(selected_rank), RANKS)][indexOfFile(File(selected_file), FILES)]
		fmt.Printf("the pieve on that square is %s \n", selected_piece)
	}

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
