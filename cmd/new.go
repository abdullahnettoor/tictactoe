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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := board.New()
		if err != nil {
			log.Panicf("Error occured:", err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
