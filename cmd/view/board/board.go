package board

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const boardSize = 3

type Board struct {
	board         [boardSize][boardSize]string
	cursor        [2]int
	player        string
	gameCompleted bool
}

func initialBoard() Board {
	return Board{
		board:  [boardSize][boardSize]string{},
		cursor: [2]int{1, 1},
		player: "X",
	}
}

func (m Board) Init() tea.Cmd { return nil }

func (m Board) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "Q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor[0] > 0 {
				m.cursor[0]--
			}
		case "down", "j":
			if m.cursor[0] < boardSize-1 {
				m.cursor[0]++
			}
		case "left", "h":
			if m.cursor[1] > 0 {
				m.cursor[1]--
			}
		case "right", "l":
			if m.cursor[1] < boardSize-1 {
				m.cursor[1]++
			}
		case "enter":
			if m.board[m.cursor[0]][m.cursor[1]] == "" {
				m.board[m.cursor[0]][m.cursor[1]] = m.player
				if checkWinner(m.board) {
					m.gameCompleted = true
					return m, nil
				}
				m.player = switchPlayer(m.player)
			}
		}
	}
	return m, nil
}

func (m Board) View() string {
	v := "'s Turn"

	if m.gameCompleted {
		v = " Won the game"
	}

	s := m.player + v

	return lipgloss.JoinVertical(lipgloss.Center,
		renderBoard(m), "\n",
		footerStyle.Render(s),
		"\n\nPress Q to exit",
	)
}

func New() error {
	p := tea.NewProgram(initialBoard())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error:", err)
		return err
	}
	return nil
}
