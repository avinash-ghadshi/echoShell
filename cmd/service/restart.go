/*
Copyright © 2024 AVINASH GHADSHI <avinashghadshi.official@gmail.com>
*/
package service

import (
	"fmt"

	"github.com/spf13/cobra"
)

// restartCmd represents the restart command
var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "restart service",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		execComand("restart", serviceName)
	},
}

func init() {
	restartCmd.Flags().StringVarP(&serviceName, "service", "s", "", "Servie name to disable")
	if err := restartCmd.MarkFlagRequired("service"); err != nil {
		fmt.Println(err.Error())
	}
}
