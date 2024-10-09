/*
Copyright © 2024 NAME HERE abdullahnettoor@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tictactoe",
	Short: "Dev friendly TicTacToe",
	Long: `Dev friendly TicTacToe
 _______    ______        ______
/_  __(_)__/_  __/__ ____/_  __/__  ___
 / / / / __// / / _ `+"`"+`/ __// / / _ \/ -_)
/_/ /_/\__//_/  \_,_/\__//_/  \___/\__/

This game is mainly focused developers as users.
Players can play tictactoe from their terminal whenever they
get bored by long hours of coding.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
