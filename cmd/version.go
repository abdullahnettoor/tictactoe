/*
Copyright © 2024 Abdullah Nettoor abdullahnettoor@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// Version holds the current version of the application
	Version = "v1.0.0"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of TicTacToe",
	Long:  `All software has versions. This is TicTacToe's.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("TicTacToe %s\n", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
