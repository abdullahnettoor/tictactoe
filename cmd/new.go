/*
Copyright © 2024 Abdullah Nettoor abdullahnettoor@gmail.com
*/
package cmd

import (
	"log"

	"github.com/abdullahnettoor/tictactoe/cmd/view/menu"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Start a new game",
	Long: `Start a new game of tictactoe.

You can choose between different game modes:
- Offline (2 Players on same terminal)
- Play against Computer
- Play Online with other players
	`,
	Run: func(cmd *cobra.Command, args []string) {
		mode, err := menu.StartMenu()
		if err != nil {
			log.Panicf("Error occurred: %v", err.Error())
		}

		err = menu.StartGame(mode)
		if err != nil {
			log.Panicf("Error occurred: %v", err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
