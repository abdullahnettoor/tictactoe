package board

import "github.com/charmbracelet/lipgloss"

var (
	xStyle      = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205"))                                              // Magenta
	oStyle      = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("39"))                                               // Cyan
	cursorStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("231")).Background(lipgloss.Color("240"))            // Highlighted cell
	emptyStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))                                                         // Grey for empty cells
	footerStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("0")).Background(lipgloss.Color("35")).Padding(0, 2) // Footer
	boardStyle  = lipgloss.NewStyle().Border(lipgloss.DoubleBorder()).Padding(1, 3)                                             // Simple border for the whole board
)

func switchPlayer(player string) string {
	if player == "X" {
		return "O"
	}
	return "X"
}

func renderBoard(m Board) string {
	var board string
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			cell := m.board[i][j]
			if cell == "" {
				cell = emptyStyle.Render(" ")
			} else {
				if cell == "X" {
					cell = xStyle.Render("X")
				} else {
					cell = oStyle.Render("O")
				}
			}

			if m.cursor[0] == i && m.cursor[1] == j {
				cell = cursorStyle.Render(cell)
			}

			if j == 0 {
				cell = " " + cell
			}
			if j == 2 {
				cell += " "
			}
			board += cell

			if j < boardSize-1 {
				board += " │ "
			}
		}

		if i < boardSize-1 {
			board += "\n───┼───┼───\n"
		}
	}
	return boardStyle.Render(board)
}

func checkWinner(board [boardSize][boardSize]string) bool {
	for i := 0; i < boardSize; i++ {
		if board[i][0] != "" && board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			return true
		}
		if board[0][i] != "" && board[0][i] == board[1][i] && board[1][i] == board[2][i] {
			return true
		}
	}
	if board[0][0] != "" && board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		return true
	}
	if board[0][2] != "" && board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		return true
	}
	return false
}
