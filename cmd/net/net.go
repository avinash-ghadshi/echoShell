/*
Copyright © 2024 AVINASH GHADSHI <avinashghadshi.official@gmail.com>
*/
package net

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// NetCmd represents the net command
var NetCmd = &cobra.Command{
	Use:   "net",
	Short: "Net package provides network functionality",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Error: No arguments provided.")
		cmd.Help() // Prints the help message for the current command
		os.Exit(0)
	},
}

func init() {

	NetCmd.AddCommand(pingCmd)
}
