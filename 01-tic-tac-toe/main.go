package main

import (
	"fmt"
)

var board [3][3]string
var currentPlayer string = "X"

var winningDirections = [8][3][2]int{
	{{0, 0}, {0, 1}, {0, 2}},
	{{1, 0}, {1, 1}, {1, 2}},
	{{2, 0}, {2, 1}, {2, 2}},
	{{0, 0}, {1, 0}, {2, 0}},
	{{0, 1}, {1, 1}, {2, 1}},
	{{0, 2}, {1, 2}, {2, 2}},
	{{0, 0}, {1, 1}, {2, 2}},
	{{0, 2}, {1, 1}, {2, 0}},
}

func main() {
	initializeBoard()
	for i := 0; i < 9; i++ {
		displayBoard()
		fmt.Printf("%s's turn (Enter row,column): ", currentPlayer)
		row, col := getMove()
		makeMove(currentPlayer, row, col)
		if checkWin() {
			fmt.Printf("%s won!", currentPlayer)
			return
		}
		switchPlayer()
	}
	fmt.Println("It's a tie!")
}

func checkWin() bool {
	for _, direction := range winningDirections {
		var didWin bool = true
		for _, cell := range direction {
			row, col := cell[0], cell[1]
			if board[row][col] != currentPlayer {
				didWin = false
				break
			}
		}
		if didWin {
			return true
		}
	}
	return false
}

func makeMove(player string, row int, col int) {
	board[row-1][col-1] = player
}

func switchPlayer() {
	if currentPlayer == "X" {
		currentPlayer = "O"
	} else {
		currentPlayer = "X"
	}
}

func getMove() (int, int) {
	var row, col int
	fmt.Scanln(&row, &col)
	if row < 1 || row > 3 || col < 1 || col > 3 {
		fmt.Println("Coordinates should be a number between 1 and 3.")
		return getMove()
	}
	if board[row-1][col-1] != " " {
		fmt.Println("This cell was already played! Choose another one.")
		return getMove()
	}
	return row, col
}

func initializeBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = " "
		}
	}
}

func displayBoard() {
	fmt.Println("   1   2   3 ")
	fmt.Println("1 ", board[0][0], "|", board[0][1], "|", board[0][2])
	fmt.Println("  ---+---+---")
	fmt.Println("2 ", board[1][0], "|", board[1][1], "|", board[1][2])
	fmt.Println("  ---+---+---")
	fmt.Println("3 ", board[2][0], "|", board[2][1], "|", board[2][2])
	fmt.Println()
}
