/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package service

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getconfCmd represents the getconf command
var getconfCmd = &cobra.Command{
	Use:   "getconf",
	Short: "Get configuration file of given service",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		//getConf(serviceName)
	},
}

func init() {
	getconfCmd.Flags().StringVarP(&serviceName, "service", "s", "", "Servie name to disable")
	if err := getconfCmd.MarkFlagRequired("service"); err != nil {
		fmt.Println(err.Error())
	}
}
