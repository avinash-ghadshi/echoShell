/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package system

import (
	"fmt"

	"github.com/spf13/cobra"
)

// findfileCmd represents the findfile command
var findfileCmd = &cobra.Command{
	Use:   "findfile",
	Short: "Find file from a directory",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		findfile(name, path)
	},
}

func init() {
	findfileCmd.Flags().StringVarP(&name, "file", "f", "", "file name to search")
	if err := findfileCmd.MarkFlagRequired("file"); err != nil {
		fmt.Println(err.Error())
	}
	findfileCmd.Flags().StringVarP(&path, "path", "p", "/", "path to search")
}
