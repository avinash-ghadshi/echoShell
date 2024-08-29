/*
Copyright © 2024 AVINASH GHADSHI <avinashghadshi.official@gmail.com>
*/
package service

import (
	"fmt"

	"github.com/spf13/cobra"
)

// unmaskCmd represents the unmask command
var unmaskCmd = &cobra.Command{
	Use:   "unmask",
	Short: "unmask service",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		execComand("unmask", serviceName)
	},
}

func init() {
	unmaskCmd.Flags().StringVarP(&serviceName, "service", "s", "", "Servie name to disable")
	if err := unmaskCmd.MarkFlagRequired("service"); err != nil {
		fmt.Println(err.Error())
	}
}
