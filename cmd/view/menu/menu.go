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
			PaddingLeft(4)

	selectedItemStyle = lipgloss.NewStyle().
				PaddingLeft(2).
				Foreground(lipgloss.Color("205")).
				SetString("❯ ")
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
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}
		case "enter":
			m.selected = true
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	s := titleStyle.Render("Select Game Mode") + "\n\n"

	for i, option := range m.options {
		if m.cursor == i {
			s += selectedItemStyle.Render(option)
		} else {
			s += itemStyle.Render(option)
		}
		s += "\n"
	}

	s += "\n" + itemStyle.Render("(Press q to quit)")
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
	switch mode {
	case LocalMultiplayer:
		return board.New()
	case LocalComputer:
		return board.NewWithComputer()
	case OnlineMultiplayer:
		return board.NewOnline()
	}
	return nil
}
