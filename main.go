package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Piece string

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

type Square string

var chessboard = [][]string{
	{"a7", "b7", "c7", "d7", "e7", "f7", "g7", "h7"},
	{"a6", "b6", "c6", "d6", "e6", "f6", "g6", "h6"},
	{"a5", "b5", "c5", "d5", "e5", "f5", "g5", "h5"},
	{"a4", "b4", "c4", "d4", "e4", "f4", "g4", "h4"},
	{"a3", "b3", "c3", "d3", "e3", "f3", "g3", "h3"},
	{"a2", "b2", "c2", "d2", "e2", "f2", "g2", "h2"},
	{"a1", "b1", "c1", "d1", "e1", "f1", "g1", "h1"},
}

type File string
type Rank string

var FILES = [8]File{"a", "b", "c", "d", "e", "f", "g", "h"}
var RANKS = [8]Rank{"1", "2", "3", "4", "5", "6", "7", "8"}

func indexOfFile(item File, array [8]File) (int, error) {
	for k, v := range array {
		if item == v {
			return k, nil
		}
	}
	return -1, errors.New("invalid file")
}

func indexOfRank(item Rank, array [8]Rank) (int, error) {
	for k, v := range array {
		if item == v {
			return k, nil
		}
	}
	return -1, errors.New("invalid rank")
}

func getPieceCoordinates(square string) ([]int, error) {
	selected_file, file_err := indexOfFile(File(strings.Split(square, "")[0]), FILES)
	selected_rank, rank_err := indexOfRank(Rank(strings.Split(square, "")[1]), RANKS)
	if file_err != nil || rank_err != nil {
		return []int{selected_file, selected_rank}, errors.New("invalid square")
	}
	return []int{selected_file, selected_rank}, nil
}

func pieceOnSquare(square string, board [][]string) (string, error) {
	selected_file := strings.Split(square, "")[0]
	selected_rank := strings.Split(square, "")[1]
	file_index, file_index_err := indexOfRank(Rank(selected_rank), RANKS)
	rank_index, rank_index_err := indexOfFile(File(selected_file), FILES)
	if file_index_err != nil || rank_index_err != nil {
		return "", errors.New("invalid square")
	}
	piece := board[7-file_index][rank_index]
	return piece, nil
}

func main() {
	FEN := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	// FEN := "rnbqkbnr/ppp2ppp/3p4/4p3/3PP3/5N2/PPP2PPP/RNBQKB1R b KQkq d3 0 3"
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

	for {
		fmt.Print("\033c")
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

		in := bufio.NewReader(os.Stdin)
		fmt.Print("Select a square:")
		user_input, err := in.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Scanf("%s", &user_input)
		fmt.Println("")

		fmt.Printf("user input => %s\n", user_input)

		selected_square, destination_square := strings.Split(user_input, " ")[0], strings.Split(user_input, " ")[1]

		selected_piece_coordinates, piece_coordinates_err := getPieceCoordinates(selected_square)
		if piece_coordinates_err != nil {
			fmt.Println(piece_coordinates_err)
		}

		destination_coordinates, piece_coordinates_err := getPieceCoordinates(destination_square)
		if piece_coordinates_err != nil {
			fmt.Println(piece_coordinates_err)
		}

		piece, piece_err := pieceOnSquare(selected_square, current_board)
		if piece_err != nil {
			fmt.Println(piece_err)
		}

		cbp := &current_board
		(*cbp)[7-destination_coordinates[1]][destination_coordinates[0]] = piece
		(*cbp)[7-selected_piece_coordinates[1]][selected_piece_coordinates[0]] = " "
	}
}

// Notes
// a1 => [0, 0]
// [0, 0] => get piece
// use piece and [0, 0] to calculate valid moves
// [x] trying to get the coordinates of a piece by giving the square "a1" and it gives back "00"
// [x] check if the input square string is a valid square on a chessboard
// - extract the FEN - Multidim slice to it's own function
// - extract render current board to its own func
// - move a piece
// - export FEN with moved square
// - start valid move logic
