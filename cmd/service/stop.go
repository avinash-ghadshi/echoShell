/*
Copyright © 2024 AVINASH GHADSHI <avinashghadshi.official@gmail.com>
*/
package service

import (
	"fmt"

	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop service",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		execComand("stop", serviceName)
	},
}

func init() {
	stopCmd.Flags().StringVarP(&serviceName, "service", "s", "", "Servie name to disable")
	if err := stopCmd.MarkFlagRequired("service"); err != nil {
		fmt.Println(err.Error())
	}
}
