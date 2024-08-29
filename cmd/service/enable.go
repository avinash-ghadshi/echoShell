/*
Copyright © 2024 AVINASH GHADSHI <avinashghadshi.official@gmail.com>
*/
package service

import (
	"fmt"

	"github.com/spf13/cobra"
)

// enableCmd represents the enable command
var enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "enable service",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		execComand("enable", serviceName)
	},
}

func init() {
	enableCmd.Flags().StringVarP(&serviceName, "service", "s", "", "Servie name to disable")
	if err := enableCmd.MarkFlagRequired("service"); err != nil {
		fmt.Println(err.Error())
	}
}
