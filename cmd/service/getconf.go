/*
Copyright © 2024 AVINASH GHADSHI <avinashghadshi.official@gmail.com>
*/
package service

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

var pattern = `\.(cf|conf|cfg|config|cnf)$`

// getconfCmd represents the getconf command
var getconfCmd = &cobra.Command{
	Use:   "getconf",
	Short: "Get configuration file of given service",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		getConf()
	},
}

func getConf() {
	confs, err := getLibraries(serviceName)
	if err != nil {
		fmt.Printf("Error retrieving configs for service '%s': %v\n", serviceName, err)
		return
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}
	sconfs := string(confs)
	fmt.Printf("Configuration files for service '%s':\n", serviceName)
	fmt.Println("--------------------------------------------")

	for _, x := range strings.Split(strings.TrimSpace(sconfs), "\n") {
		if re.MatchString(x) {
			fmt.Println(x)
		}
	}
	fmt.Println("--------------------------------------------")
}

func init() {
	getconfCmd.Flags().StringVarP(&serviceName, "service", "s", "", "Servie name to disable")
	if err := getconfCmd.MarkFlagRequired("service"); err != nil {
		fmt.Println(err.Error())
	}
}
