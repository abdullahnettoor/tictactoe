# 🎮 Terminal TicTacToe

> A developer-friendly TicTacToe game that runs in your terminal. Take a break from coding and enjoy a quick game without leaving your development environment! ⌨️

```ascii
 _______    ______        ______
/_  __(_)__/_  __/__  ___/_  __/__  ___
 / / / / __// / / _ `/ __// / / _ \/ -_)
/_/ /_/\__//_/  \_,_/\__//_/  \___/\__/
```

<div align="center">

[![Go Version](https://img.shields.io/badge/Go-1.21.7+-00ADD8?style=flat-square&logo=go)](https://golang.org)
[![Version](https://img.shields.io/badge/Version-1.0.0-blue.svg?style=flat-square)](https://github.com/abdullahnettoor/tictactoe/releases)
[![License](https://img.shields.io/badge/License-MIT-green.svg?style=flat-square)](LICENSE)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)

</div>

## ✨ Features

- 🎯 Play TicTacToe directly in your terminal
- 🤖 Challenge the unbeatable AI opponent
- ⌨️ Vim-style navigation (h,j,k,l) or arrow keys
- 🎨 Colorful interface with intuitive controls
- 👥 Multiple game modes (vs Human, vs Computer)
- 🎭 Simple and clean UI with borders and highlighting
- 🔄 Play again feature after game completion

## 🚀 Installation

1. Make sure you have Go installed (version 1.21.7 or higher)
2. Install the game using:

```bash
go install github.com/abdullahnettoor/tictactoe@v1.0.0
```

## 🎮 How to Play

1. Start a new game:
```bash
tictactoe new
```

2. Select game mode:
   - 👥 Local Multiplayer (2 Players)
   - 🤖 Local Computer (vs AI)
   - 🌐 Online Multiplayer (Coming soon!)

3. Controls:
   - 🕹️ Move: Arrow keys or Vim keys (h,j,k,l)
   - ✅ Place mark: Enter
   - 🔄 Play again: P
   - ❌ Quit game: Q or Ctrl+C

4. Game Rules:
   - Players take turns placing their marks (X and O)
   - First player to get three in a row wins!
   - In Computer mode, try to beat the AI (if you can! 😉)

## 💻 Development

### Prerequisites

- 🔧 Go 1.21.7 or higher
- 📦 Dependencies:
  - [github.com/charmbracelet/bubbletea](https://github.com/charmbracelet/bubbletea) - Terminal UI framework
  - [github.com/charmbracelet/lipgloss](https://github.com/charmbracelet/lipgloss) - Style definitions
  - [github.com/spf13/cobra](https://github.com/spf13/cobra) - CLI framework

### 🛠️ Building from Source

1. Clone the repository:
```bash
git clone https://github.com/abdullahnettoor/tictactoe.git
```

2. Navigate to the project directory:
```bash
cd tictactoe
```

3. Build and run:
```bash
make run
```

## 📁 Project Structure

```
.
├── 📄 main.go              # Entry point of the application
├── 📂 cmd/
│   ├── 📄 root.go         # Root command configuration
│   ├── 📄 new.go          # New game command handler
│   └── 📂 view/
│       ├── 📂 board/      # Game board implementation
│       │   ├── 📄 board.go # Core game logic and UI
│       │   ├── 📄 ai.go   # AI player implementation
│       │   └── 📄 utils.go # Helper functions and styling
│       └── 📂 menu/       # Game mode selection menu
└── 📄 go.mod              # Go module file
```

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👤 Author

Made with ❤️ by [Abdullah Nettoor](https://github.com/abdullahnettoor)

