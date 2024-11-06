package menu

import (
	"github.com/abdullahnettoor/tictactoe/cmd/view/board"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type GameMode int

const (
	LocalMultiplayer GameMode = iota
	LocalComputer
	OnlineMultiplayer
)

type model struct {
	cursor   int
	options  []string
	selected bool
}

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("205")).
			MarginLeft(2)

	itemStyle = lipgloss.NewStyle().
			PaddingLeft(2)

	selectedItemStyle = lipgloss.NewStyle().
				PaddingLeft(2).
				Foreground(lipgloss.Color("205")).
				Background(lipgloss.Color("55")).
				Bold(true)
)

func initialModel() model {
	return model{
		options: []string{
			"Local Multiplayer (2 Players)",
			"Local Computer (vs AI)",
			"Online Multiplayer",
		},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(m.options) - 2
			}
		case "down", "j":
			m.cursor++
			if m.cursor > len(m.options)-2 {
				m.cursor = 0
			}
		case "enter":
			if m.cursor < len(m.options)-1 {
				m.selected = true
				return m, tea.Quit
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	s := titleStyle.Render("Select Game Mode")
	s += "\n\n"

	for i, option := range m.options {
		var renderedOption string
		if m.cursor == i {
			renderedOption = selectedItemStyle.Render("❯ " + option)
		} else {
			renderedOption = itemStyle.Render("  " + option)
		}

		if i == int(OnlineMultiplayer) {
			s += renderedOption + " (Coming Soon!)\n"
			continue
		}

		s += renderedOption + "\n"
	}

	return s
}

func StartMenu() (GameMode, error) {
	m := initialModel()
	p := tea.NewProgram(m)
	finalModel, err := p.Run()
	if err != nil {
		return 0, err
	}

	if !finalModel.(model).selected {
		return 0, nil
	}

	return GameMode(finalModel.(model).cursor), nil
}

func StartGame(mode GameMode) error {
	var gameMode string
	var isComputer bool

	switch mode {
	case LocalMultiplayer:
		gameMode = "Local Multiplayer"
		isComputer = false
	case LocalComputer:
		gameMode = "Playing Against Computer"
		isComputer = true
	case OnlineMultiplayer:
		gameMode = "Online Multiplayer"
		isComputer = false
	}

	p := tea.NewProgram(board.NewBoard(isComputer, gameMode))
	_, err := p.Run()
	return err
}
