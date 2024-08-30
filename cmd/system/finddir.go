/*
Copyright © 2024 AVINASH GHADSHI <avinashghadshi.official@gmail.com>
*/
package system

import (
	"fmt"

	"github.com/spf13/cobra"
)

// finddirCmd represents the finddir command
var finddirCmd = &cobra.Command{
	Use:   "finddir",
	Short: "find directory in specified path",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		finddir(name, path)
	},
}

func init() {
	finddirCmd.Flags().StringVarP(&name, "dir", "d", "", "directory to search")
	if err := finddirCmd.MarkFlagRequired("dir"); err != nil {
		fmt.Println(err.Error())
	}
	finddirCmd.Flags().StringVarP(&path, "path", "p", "/", "path to search")
}
