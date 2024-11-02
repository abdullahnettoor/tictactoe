package board

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const boardSize = 3

type Board struct {
	board          [boardSize][boardSize]string
	cursor         [2]int
	player         string
	gameCompleted  bool
	isComputerMode bool
	gameMode       string
}

func NewBoard(isComputer bool, mode string) Board {
	return Board{
		board:          [boardSize][boardSize]string{},
		cursor:         [2]int{1, 1},
		player:         "X",
		isComputerMode: isComputer,
		gameMode:       mode,
	}
}

func (m Board) Init() tea.Cmd { return nil }

func (m Board) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "Q":
			return m, tea.Quit

		case "p", "P":
			if m.gameCompleted {
				// Start a new game with same mode
				return NewBoard(m.isComputerMode, m.gameMode), nil
			}

		case "enter":
			if !m.gameCompleted && m.board[m.cursor[0]][m.cursor[1]] == "" {
				m.board[m.cursor[0]][m.cursor[1]] = m.player
				if winner := checkWinner(m.board); winner != "" {
					m.gameCompleted = true
					return m, nil
				}

				if isBoardFull(m.board) {
					m.gameCompleted = true
					return m, nil
				}

				m.player = switchPlayer(m.player)

				// Computer move
				if m.isComputerMode && m.player == "O" {
					move := getBestMove(m.board)
					m.board[move[0]][move[1]] = m.player
					if winner := checkWinner(m.board); winner != "" {
						m.gameCompleted = true
						return m, nil
					}
					m.player = switchPlayer(m.player)
				}
			}

		case "up", "k":
			if !m.gameCompleted {
				if m.cursor[0] > 0 {
					m.cursor[0]--
				}
			}
		case "down", "j":
			if !m.gameCompleted {
				if m.cursor[0] < boardSize-1 {
					m.cursor[0]++
				}
			}
		case "left", "h":
			if !m.gameCompleted {
				if m.cursor[1] > 0 {
					m.cursor[1]--
				}
			}
		case "right", "l":
			if !m.gameCompleted {
				if m.cursor[1] < boardSize-1 {
					m.cursor[1]++
				}
			}
		}
	}
	return m, nil
}

func (m Board) View() string {
	var s string

	if m.gameCompleted {
		winner := checkWinner(m.board)
		if winner != "" {
			if m.isComputerMode {
				if winner == aiPlayer {
					s = "You lost to Computer!"
				} else {
					s = "You won against Computer!"
				}
			} else {
				s = winner + " Won the game!"
			}
		} else {
			s = "Game Draw!"
		}
		s += "\n\nPress P to play again"
	} else {
		if m.isComputerMode && m.player == aiPlayer {
			s = "Computer's Turn"
		} else {
			s = m.player + "'s Turn"
		}
	}

	return lipgloss.JoinVertical(lipgloss.Center,
		titleStyle.Render(m.gameMode),
		"\n",
		renderBoard(m),
		"\n",
		footerStyle.Render(s),
		"\n\nPress Q to exit",
	)
}

func New() error {
	p := tea.NewProgram(NewBoard(false, "Human vs Human"))
	if _, err := p.Run(); err != nil {
		fmt.Println("Error:", err)
		return err
	}
	return nil
}

func NewWithComputer() error {
	p := tea.NewProgram(NewBoard(true, "Human vs Computer"))
	if _, err := p.Run(); err != nil {
		fmt.Println("Error:", err)
		return err
	}
	return nil
}

// TODO
func NewOnline() error {
	return fmt.Errorf("online mode not implemented yet")
}
