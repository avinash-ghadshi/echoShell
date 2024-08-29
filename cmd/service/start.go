/*
Copyright © 2024 AVINASH GHADSHI <avinashghadshi.official@gmail.com>
*/
package service

import (
	"fmt"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start service",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		execComand("start", serviceName)
	},
}

func init() {
	startCmd.Flags().StringVarP(&serviceName, "service", "s", "", "Servie name to disable")
	if err := startCmd.MarkFlagRequired("service"); err != nil {
		fmt.Println(err.Error())
	}
}
