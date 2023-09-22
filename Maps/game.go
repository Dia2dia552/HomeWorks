package main

import (
	"fmt"
)

var board = make(map[int]string)
var currentPlayer = "X"
var gameOver = false

func main() {
	initializeBoard()
	displayBoard()

	for !gameOver {
		makeMove()
		displayBoard()
		checkWinner()
		switchPlayer()
	}
}

func initializeBoard() {
	for i := 1; i <= 9; i++ {
		board[i] = " "
	}
}

func displayBoard() {
	fmt.Println(" 1 | 2 | 3 ")
	fmt.Println("---+---+---")
	fmt.Println(" 4 | 5 | 6 ")
	fmt.Println("---+---+---")
	fmt.Println(" 7 | 8 | 9 ")
	fmt.Println()

	fmt.Println("Current Player:", currentPlayer)
	fmt.Println("-------------")
	fmt.Println(" " + board[1] + " | " + board[2] + " | " + board[3] + " ")
	fmt.Println("---+---+---")
	fmt.Println(" " + board[4] + " | " + board[5] + " | " + board[6] + " ")
	fmt.Println("---+---+---")
	fmt.Println(" " + board[7] + " | " + board[8] + " | " + board[9] + " ")
}

func makeMove() {
	fmt.Print("Enter a position (1-9) to place your symbol: ")
	var position int
	_, err := fmt.Scanln(&position)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number between 1 and 9.")
		return
	}

	if position < 1 || position > 9 || board[position] != " " {
		fmt.Println("Invalid move. Try again.")
		return
	}

	board[position] = currentPlayer
}

func switchPlayer() {
	if currentPlayer == "X" {
		currentPlayer = "O"
	} else {
		currentPlayer = "X"
	}
}

func checkWinner() {
	if checkLine(1, 2, 3) || checkLine(4, 5, 6) || checkLine(7, 8, 9) ||
		checkLine(1, 4, 7) || checkLine(2, 5, 8) || checkLine(3, 6, 9) ||
		checkLine(1, 5, 9) || checkLine(3, 5, 7) {
		displayBoard()
		fmt.Println("Player", currentPlayer, "wins!")
		gameOver = true
	} else if boardIsFull() {
		displayBoard()
		fmt.Println("It's a draw!")
		gameOver = true
	}
}

func checkLine(a, b, c int) bool {
	return board[a] != " " && board[a] == board[b] && board[b] == board[c]
}

func boardIsFull() bool {
	for i := 1; i <= 9; i++ {
		if board[i] == " " {
			return false
		}
	}
	return true
}
