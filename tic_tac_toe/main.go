// Simple CLI tic tac toe game 
//

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	// Game loop for multiple rounds
	for {
		playGame()
		
		fmt.Print("\nPlay again? (y/n): ")
		reader := bufio.NewReader(os.Stdin)
		playAgain, _ := reader.ReadString('\n')
		playAgain = strings.TrimSpace(strings.ToLower(playAgain))
		
		if playAgain != "y" && playAgain != "yes" {
			fmt.Println("Thanks for playing!")
			break
		}
	}
}

func playGame() {
	// Step 2: Initialize the board (1-9 positions)
	board := [3][3]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}
	currentPlayer := "X"
	moves := 0

	// Step 8: Main game loop
	for {
		clearScreen()
		displayBoard(board) // Step 3
		
		// Step 4: Get player input
		fmt.Printf("\nPlayer %s, choose a cell (1-9): ", currentPlayer)
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		
		position, err := strconv.Atoi(input)
		if err != nil || position < 1 || position > 9 {
			fmt.Println("Invalid input! Please enter a number between 1-9.")
			pressEnterToContinue()
			continue
		}
		
		// Step 5: Check if cell is available and place mark
		row, col := getCoordinates(position)
		if board[row][col] == "X" || board[row][col] == "O" {
			fmt.Println("Cell already taken! Choose another.")
			pressEnterToContinue()
			continue
		}
		
		// Place the mark
		board[row][col] = currentPlayer
		moves++
		
		// Step 6: Check win condition
		if checkWin(board, currentPlayer) {
			clearScreen()
			displayBoard(board)
			fmt.Printf("\n🎉 Player %s wins! 🎉\n", currentPlayer)
			return
		}
		
		// Step 7: Check for draw
		if moves == 9 {
			clearScreen()
			displayBoard(board)
			fmt.Println("\n🤝 It's a draw! 🤝")
			return
		}
		
		// Switch players
		if currentPlayer == "X" {
			currentPlayer = "O"
		} else {
			currentPlayer = "X"
		}
	}
}

// Step 3: Display the board
func displayBoard(board [3][3]string) {
	fmt.Println("\n─────────────")
	for i := 0; i < 3; i++ {
		fmt.Printf("│ %s │ %s │ %s │\n", board[i][0], board[i][1], board[i][2])
		fmt.Println("─────────────")
	}
}

// Convert 1-9 position to row, col coordinates
func getCoordinates(pos int) (int, int) {
	pos-- // Convert to 0-based index
	return pos / 3, pos % 3
}

// Step 6: Check win conditions
func checkWin(board [3][3]string, player string) bool {
	// Check rows
	for i := 0; i < 3; i++ {
		if board[i][0] == player && board[i][1] == player && board[i][2] == player {
			return true
		}
	}
	
	// Check columns
	for j := 0; j < 3; j++ {
		if board[0][j] == player && board[1][j] == player && board[2][j] == player {
			return true
		}
	}
	
	// Check diagonals
	if board[0][0] == player && board[1][1] == player && board[2][2] == player {
		return true
	}
	if board[0][2] == player && board[1][1] == player && board[2][0] == player {
		return true
	}
	
	return false
}

// Step 9: Clear screen for better readability
func clearScreen() {
	var cmd *exec.Cmd
	
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Helper to pause and continue
func pressEnterToContinue() {
	fmt.Print("\nPress Enter to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
