package game

import (
	"fmt"
	"os"
	"strconv"
	"tictactoe/winner"

	"github.com/fatih/color"
)

type squareValue = byte

const (
	EMPTY_SQUARE squareValue = ' '
	X_SQAURE     squareValue = 'X'
	O_SQAURE     squareValue = 'O'
)

type game struct {
	isGameActive bool

	// gameState
	board  [9]byte
	winner int

	// 'X' or 'O'
	currentMove byte
}

func New() game {
	var filledBoard [9]byte

	for i := range filledBoard {
		filledBoard[i] = EMPTY_SQUARE
	}

	g := game{isGameActive: true, board: filledBoard, winner: winner.NONE, currentMove: 'X'}
	return g
}

func IsActive(g *game) bool {
	return g.isGameActive
}

func GetColorSquare(ch byte) string {
	if ch == X_SQAURE {
		return color.New(color.Bold, color.FgRed).SprintFunc()(string(ch))
	} else {
		return color.New(color.Bold, color.FgBlue).SprintFunc()(string(ch))
	}
}

func convCharToWinner(ch byte) int {
	if ch == X_SQAURE {
		return winner.X
	} else {
		return winner.O
	}
}

func CheckWinner(g *game) int {
	// check columns
	for i := 0; i < 3; i++ {
		if g.board[i] == g.board[i+3] && g.board[i] == g.board[i+6] {
			if g.board[i] != ' ' {
				return convCharToWinner(g.board[i])
			}
		}
	}

	// check rows
	for i := 0; i < 9; i += 3 {
		if g.board[i] == g.board[i+1] && g.board[i] == g.board[i+2] {
			if g.board[i] != ' ' {
				return convCharToWinner(g.board[i])
			}
		}
	}

	// check diagonals
	if g.board[0] == g.board[4] && g.board[0] == g.board[8] {
		if g.board[0] != ' ' {
			return convCharToWinner(g.board[0])
		}

	}

	if g.board[2] == g.board[4] && g.board[2] == g.board[6] {
		if g.board[2] != ' ' {
			return convCharToWinner(g.board[2])
		}

	}

	// check for empty squares
	hasEmpty := false
	for _, val := range g.board {
		if val == ' ' {
			hasEmpty = true
			break
		}
	}

	if !hasEmpty {
		return winner.DRAW
	}

	return winner.NONE
}

func UpdateGame(g *game, inputPos int) {
	if g.currentMove == 'X' {
		g.board[inputPos-1] = X_SQAURE
	} else {
		g.board[inputPos-1] = O_SQAURE
	}

	winCheckRes := CheckWinner(g)

	if winCheckRes != winner.NONE {
		g.winner = winCheckRes
		g.isGameActive = false
		return
	}

	if g.currentMove == 'X' {
		g.currentMove = 'O'
	} else {
		g.currentMove = 'X'
	}
}

func PrintBoard(g *game) {
	// clear screen
	fmt.Print("\033[H\033[2J")
	fmt.Println(" _____________")

	var msg string

	if g.isGameActive {
		msg = fmt.Sprintf("%s gājiens", GetColorSquare(g.currentMove))
	} else {
		switch g.winner {
		case winner.DRAW:
			msg = "Neizšķirts!"
		case winner.X:
			msg = fmt.Sprintf("%s uzvar!", GetColorSquare(X_SQAURE))
		case winner.O:
			msg = fmt.Sprintf("%s uzvar!", GetColorSquare(O_SQAURE))
		default:
			msg = ""
		}
	}

	for index, val := range g.board {
		var squareChar string

		if val == ' ' {
			if g.isGameActive {
				squareChar = fmt.Sprintf("%c", index+49)
			} else {
				squareChar = " "
			}
		} else {
			squareChar = GetColorSquare(val)
		}

		switch index {
		case 0:
			fmt.Printf("/  %s │", squareChar)
		case 3:
			fmt.Printf("|  %s │", squareChar)
		case 6:
			fmt.Printf("\\  %s │", squareChar)
		case 1:
			fallthrough
		case 4:
			fallthrough
		case 7:
			fmt.Printf(" %s │", squareChar)
		case 2:
			fmt.Printf(" %s  \\\n", squareChar)
			fmt.Println("| ───┼───┼─── |")
		case 5:
			fmt.Printf(" %s  |    %s\n", squareChar, msg)
			fmt.Println("| ───┼───┼─── |")
		case 8:
			fmt.Printf(" %s  /\n", squareChar)
		}
	}

	fmt.Println(" -------------")
	fmt.Println("        \\   ^__^")
	fmt.Println("         \\  (oo)\\_______")
	fmt.Println("            (__)\\       )\\/\\")
	fmt.Println("               ||----w |")
	fmt.Println("               ||     ||")

	fmt.Print("\n")
}

func GetNextMoveInput(g *game) int {
	fmt.Print("Ievadi skaitli (1-9): ")

	var input string

	fmt.Scanln(&input)

	if input == "exit" {
		os.Exit(0)
	}

	num, err := strconv.Atoi(input)

	if err != nil || num < 1 || num > 9 || g.board[num-1] != ' ' {
		PrintBoard(g)
		return GetNextMoveInput(g)
	}

	return num
}
