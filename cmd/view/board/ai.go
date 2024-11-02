package board

// Score values for Minimax algorithm
const (
	aiWin      = 10
	humanWin   = -10
	draw       = 0
	aiPlayer   = "O"
	aiOpponent = "X"
)

func getBestMove(board [boardSize][boardSize]string) [2]int {
	var bestMove [2]int
	bestScore := -1000

	// Try all possible moves
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if board[i][j] == "" {
				// Make the move
				board[i][j] = aiPlayer
				// Get score for this move
				score := minimax(board, 0, false)
				// Undo the move
				board[i][j] = ""

				// Update best move if current score is better
				if score > bestScore {
					bestScore = score
					bestMove = [2]int{i, j}
				}
			}
		}
	}
	return bestMove
}

func minimax(board [boardSize][boardSize]string, depth int, isMaximizing bool) int {
	if winner := checkWinner(board); winner != "" {
		if winner == aiPlayer {
			return aiWin
		}
		if winner == aiOpponent {
			return humanWin
		}
	}

	if isBoardFull(board) {
		return draw
	}

	if isMaximizing {
		bestScore := -1000
		for i := 0; i < boardSize; i++ {
			for j := 0; j < boardSize; j++ {
				if board[i][j] == "" {
					board[i][j] = aiPlayer
					score := minimax(board, depth+1, false)
					board[i][j] = ""
					bestScore = max(score, bestScore)
				}
			}
		}
		return bestScore
	}

	bestScore := 1000
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if board[i][j] == "" {
				board[i][j] = aiOpponent
				score := minimax(board, depth+1, true)
				board[i][j] = ""
				bestScore = min(score, bestScore)
			}
		}
	}
	return bestScore
}

func isBoardFull(board [boardSize][boardSize]string) bool {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if board[i][j] == "" {
				return false
			}
		}
	}
	return true
}
