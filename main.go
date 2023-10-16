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
	"i": "\u2022",
	"x": "\u2717",
	"#": "\u2592",
}

type Square string

type (
	File string
	Rank string
)

var (
	FILES = [8]File{"a", "b", "c", "d", "e", "f", "g", "h"}
	RANKS = [8]Rank{"1", "2", "3", "4", "5", "6", "7", "8"}
)

var CHESSBOARD = [][]string{
	{"h1", "h2", "h3", "h4", "h5", "h6", "h7", "h8"},
	{"g1", "g2", "g3", "g4", "g5", "g6", "g7", "g8"},
	{"f1", "f2", "f3", "f4", "f5", "f6", "f7", "f8"},
	{"e1", "e2", "e3", "e4", "e5", "e6", "e7", "e8"},
	{"d1", "d2", "d3", "d4", "d5", "d6", "d7", "d8"},
	{"c1", "c2", "c3", "c4", "c5", "c6", "c7", "c8"},
	{"b1", "b2", "b3", "b4", "b5", "b6", "b7", "b8"},
	{"a1", "a2", "a3", "a4", "a5", "a6", "a7", "a8"},
}

func IndexOfFile(item File, array [8]File) (int, error) {
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
	fmt.Println("selected_square", square)
	square = strings.TrimSpace(square)
	selected_file, file_err := IndexOfFile(File(strings.Split(square, "")[0]), FILES)
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
	rank_index, rank_index_err := IndexOfFile(File(selected_file), FILES)
	if file_index_err != nil || rank_index_err != nil {
		return "", errors.New("invalid square")
	}
	piece := board[7-file_index][rank_index]
	return piece, nil
}

func getSquareOnboard(coordinates []int, board *[][]string) string {
	bp := *board
	piece := bp[7-coordinates[1]][coordinates[0]]
	return piece
}

func getForwardSquare(coordinates []int, board *[][]string) string {
	bp := *board
	piece := bp[7-coordinates[1]-1][coordinates[0]]
	return piece
}

func getDoubleForwardSquare(coordinates []int, board *[][]string) string {
	bp := *board
	piece := bp[7-coordinates[1]-2][coordinates[0]]
	return piece
}

func getForwardDiagonalLeftSquare(coordinates []int, board *[][]string) string {
	bp := *board
	piece := bp[7-coordinates[1]-1][coordinates[0]+1]
	return piece
}

func getForwardDiagonalRightSquare(coordinates []int, board *[][]string) string {
	bp := *board
	piece := bp[7-coordinates[1]-1][coordinates[0]+1]
	return piece
}

func movePiece(selected_square []int, destination_square []int, board *[][]string) {
	cbp := *board
	(cbp)[7-destination_square[1]][destination_square[0]] = getSquareOnboard(selected_square, board)
	(cbp)[7-selected_square[1]][selected_square[0]] = " "
}

func pawnMove(
	piece_coordinates []int,
	destination_coordinates []int,
	board *[][]string,
) (string, error) {
	// can't move above 2 squares ahead
	if destination_coordinates[1]-piece_coordinates[1] > 2 {
		return "", errors.New("invalid move")
	}

	// 2nd Rank move
	if piece_coordinates[1] == 1 {
		if getForwardSquare(piece_coordinates, board) != " " ||
			getDoubleForwardSquare(piece_coordinates, board) != " " {
			return "", errors.New("invalid move")
		}
		movePiece(piece_coordinates, destination_coordinates, board)
		return "ok", nil
	}

	if destination_coordinates[1]-piece_coordinates[1] > 1 {
		return "", errors.New("invalid move")
	}

	diagonal_right_square_piece := getForwardDiagonalRightSquare(piece_coordinates, board)
	diagonal_right_square_coordinates, err := getPieceCoordinates(diagonal_right_square_piece)
	if err != nil {
		fmt.Println(err)
	}
	diagonal_left_square_piece := getForwardDiagonalLeftSquare(piece_coordinates, board)
	diagonal_left_square_coordinates, err := getPieceCoordinates(diagonal_right_square_piece)
	if err != nil {
		fmt.Println(err)
	}
	forward_square_piece := getForwardSquare(piece_coordinates, board)
	forward_square_coordinates, err := getPieceCoordinates(forward_square_piece)
	if err != nil {
		fmt.Println(err)
	}

	// capture diagonal right
	if destination_coordinates[0] == diagonal_right_square_coordinates[0] &&
		destination_coordinates[1] == diagonal_right_square_coordinates[1] &&
		diagonal_right_square_piece != " " {
		fmt.Println("capture diagonal right")
		movePiece(
			piece_coordinates,
			[]int{destination_coordinates[0] + 1, destination_coordinates[1]},
			board,
		)
		return "ok", nil
	}

	// capture diagonal left
	if destination_coordinates[0] == diagonal_left_square_coordinates[0] &&
		destination_coordinates[1] == diagonal_left_square_coordinates[1] &&
		diagonal_left_square_piece != " " {
		fmt.Println("capture diagonal left")
		movePiece(
			piece_coordinates,
			[]int{destination_coordinates[0], destination_coordinates[1] - 1},
			board,
		)
		return "ok", nil
	}

	if destination_coordinates[0] != piece_coordinates[0] {
		return "", errors.New("invalid move")
	}

	// Normal forward move
	if forward_square_coordinates[0] == destination_coordinates[0] &&
		forward_square_coordinates[1] == destination_coordinates[1] &&
		getForwardSquare(piece_coordinates, board) != " " {
		return "", errors.New("invalid move")
	}

	// promote
	if destination_coordinates[1] == 7 {
		in := bufio.NewReader(os.Stdin)
		fmt.Print("")
		promotion_piece, err := in.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		promotion_piece = strings.TrimSpace(promotion_piece)
		if promotion_piece == "Q" || promotion_piece == "B" || promotion_piece == "N" ||
			promotion_piece == "R" {
			cbp := *board
			(cbp)[7-destination_coordinates[1]][destination_coordinates[0]] = promotion_piece
		}
		return "ok", nil
	}

	movePiece(piece_coordinates, destination_coordinates, board)
	return "ok", nil
}

func initialiseBoard(FEN_string *string) *[][]string {
	FEN := FEN_string
	FEN_layout := strings.Split(strings.Split(*FEN, " ")[0], "")

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
				// current_board[row][col] = display[Piece(FEN_piece)]
				current_board[row][col] = FEN_piece
				col++
			}
		}
	}

	return &current_board
}

func startGameLoop(board *[][]string) {
	for {
		current_board := *board
		// fmt.Print("\033c")
		fmt.Println("   a b c d e f g h")
		fmt.Println(" ╔═════════════════╗")
		for i := 0; i < 8; i++ {
			fmt.Print(8 - i)
			fmt.Print("║ ")
			for j := 0; j < 8; j++ {
				fmt.Print(current_board[i][j])
				fmt.Print(" ")
			}
			fmt.Print("║")
			fmt.Println()
		}
		fmt.Println(" ╚═════════════════╝")

		in := bufio.NewReader(os.Stdin)
		fmt.Print("move (eg. e2 e4):")
		user_input, err := in.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		selected_square, destination_square := strings.Split(user_input, " ")[0], strings.Split(user_input, " ")[1]

		selected_piece_coordinates, piece_coordinates_err := getPieceCoordinates(selected_square)
		if piece_coordinates_err != nil {
			fmt.Println(piece_coordinates_err)
		}

		destination_coordinates, destination_coordinates_err := getPieceCoordinates(
			destination_square,
		)
		if destination_coordinates_err != nil {
			fmt.Println(destination_coordinates_err)
		}

		piece, piece_err := pieceOnSquare(selected_square, current_board)
		if piece_err != nil {
			fmt.Println(piece_err)
		}

		// move piece

		if piece == "p" || piece == "P" {
			_, err := pawnMove(selected_piece_coordinates, destination_coordinates, board)
			if err != nil {
				fmt.Println(err)
			}
		}

		// cbp := &current_board
		// (*cbp)[7-destination_coordinates[1]][destination_coordinates[0]] = piece
		// (*cbp)[7-selected_piece_coordinates[1]][selected_piece_coordinates[0]] = " "
	}
}

func startNewGame(FEN *string) {
	*FEN = "rnbqkbnr/p1pppppp/8/1p6/8/N7/PPPPPPPP/R1BQKBNR w KQkq - 0 1"
	board := initialiseBoard(FEN)
	startGameLoop(board)
}

func main() {
	invalid_input := false
	for {
		var FEN string
		// fmt.Print("\033c")
		fmt.Println("╔═══════════════════╗")
		fmt.Println("║       Chess       ║")
		fmt.Println("║                   ║")
		fmt.Println("║ 1. New Game       ║")
		// fmt.Println("║ 2. Load Game      ║")
		// fmt.Println("║ 3. Instructions   ║")
		fmt.Println("║ 2. Quit           ║")
		fmt.Println("║                   ║")
		if invalid_input {
			fmt.Println("║   invalid input   ║")
		} else {
			fmt.Println("║                   ║")
		}
		fmt.Println("╚═══════════════════╝")

		in := bufio.NewReader(os.Stdin)
		fmt.Print("")
		menu_input, err := in.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		menu_input = strings.TrimSpace(menu_input)

		if menu_input == "1" {
			startNewGame(&FEN)
		} else if menu_input == "2" {
			fmt.Println("Quitting the game...")
			break // Exit the loop and end the program
		} else {
			invalid_input = true
		}
	}
}
