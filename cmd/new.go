/*
Copyright © 2024 Abdullah Nettoor abdullahnettoor@gmail.com
*/
package cmd

import (
	"log"

	"github.com/abdullahnettoor/tictactoe/cmd/view/board"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Start a new game",
	Long: `Start a new game of tictactoe.

The game can be played by two players in the same terminal window.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		err := board.New()
		if err != nil {
			log.Panicf("Error occured: %v", err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
