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
[![License](https://img.shields.io/badge/License-MIT-green.svg?style=flat-square)](LICENSE)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)

</div>

## ✨ Features

- 🎯 Play TicTacToe directly in your terminal
- ⌨️ Vim-style navigation (h,j,k,l) or arrow keys
- 🎨 Colorful interface with intuitive controls
- 👥 Two-player gameplay on the same terminal
- 🎭 Simple and clean UI with borders and highlighting

## 🚀 Installation

1. Make sure you have Go installed (version 1.21.7 or higher)
2. Install the game using:

```bash
go install github.com/abdullahnettoor/tictactoe@latest
```

## 🎮 How to Play

1. Start a new game:
```bash
tictactoe new
```

2. Controls:
   - 🕹️ Move: Arrow keys or Vim keys (h,j,k,l)
   - ✅ Place mark: Enter
   - ❌ Quit game: q or Ctrl+C

3. Players take turns placing their marks (X and O)
4. First player to get three in a row (horizontal, vertical, or diagonal) wins! 🏆

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

3. Build the project:
```bash
go build
```

## 📁 Project Structure

```
.
├── 📄 main.go              # Entry point of the application
├── 📂 cmd/
│   ├── 📄 root.go         # Root command configuration
│   ├── 📄 new.go          # New game command handler
│   └── 📂 view/board/     # Game board implementation
│       ├── 📄 board.go    # Core game logic and UI
│       └── 📄 utils.go    # Helper functions and styling
└── 📄 go.mod              # Go module file
```

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👤 Author

Made with ❤️ by [Abdullah Nettoor](https://github.com/abdullahnettoor) 

