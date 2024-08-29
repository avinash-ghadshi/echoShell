/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package service

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getservicefileCmd represents the getservicefile command
var getservicefileCmd = &cobra.Command{
	Use:   "getservicefile",
	Short: "get service file of a given service",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		getService(serviceName)
	},
}

func init() {
	getservicefileCmd.Flags().StringVarP(&serviceName, "service", "s", "", "Servie name to disable")
	if err := getservicefileCmd.MarkFlagRequired("service"); err != nil {
		fmt.Println(err.Error())
	}
}
