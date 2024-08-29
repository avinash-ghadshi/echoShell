/*
Copyright © 2024 AVINASH GHADSHI <avinashghadshi.official@gmail.com>
*/
package service

import (
	"fmt"

	"github.com/spf13/cobra"
)

// disableCmd represents the disable command
var disableCmd = &cobra.Command{
	Use:   "disable",
	Short: "disable service",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		execComand("disable", serviceName)
	},
}

func init() {
	disableCmd.Flags().StringVarP(&serviceName, "service", "s", "", "Servie name to disable")
	if err := disableCmd.MarkFlagRequired("service"); err != nil {
		fmt.Println(err.Error())
	}
}
