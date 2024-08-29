/*
Copyright © 2024 AVINASH GHADSHI <avinashghadshi.official@gmail.com>
*/
package service

import (
	"fmt"

	"github.com/spf13/cobra"
)

// maskCmd represents the mask command
var maskCmd = &cobra.Command{
	Use:   "mask",
	Short: "mask service",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		execComand("mask", serviceName)
	},
}

func init() {
	maskCmd.Flags().StringVarP(&serviceName, "service", "s", "", "Servie name to disable")
	if err := maskCmd.MarkFlagRequired("service"); err != nil {
		fmt.Println(err.Error())
	}
}
