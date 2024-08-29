/*
Copyright © 2024 AVINASH GHADSHI <avinashghadshi.official@gmail.com>
*/
package service

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var serviceName string

// ServiceCmd represents the service command
var ServiceCmd = &cobra.Command{
	Use:   "service",
	Short: "service package contains commands to manage services",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Error: No arguments provided.")
		cmd.Help() // Prints the help message for the current command
		os.Exit(0)
	},
}

func serviceExists(name string) bool {
	cmd := exec.Command("systemctl", "list-units", "--type=service", "--all", "--no-pager")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error running systemctl command: %v\n", err)
		return false
	}
	return bytes.Contains(out.Bytes(), []byte(name))
}

func execComand(action, sn string) {
	if !serviceExists(sn) {
		fmt.Printf("Service '%s' not found\n", sn)
		return
	}
	cmd := exec.Command("systemctl", action, sn)
	output, _ := cmd.CombinedOutput()
	fmt.Println(string(output))
}

func getService(sn string) {
	if !serviceExists(sn) {
		fmt.Printf("Service '%s' not found\n", sn)
		return
	}
	cmd := exec.Command("systemctl", "show", sn, "--property=FragmentPath")
	output, _ := cmd.CombinedOutput()
	aOutput := strings.Split(string(output), "=")
	if len(aOutput) != 2 {
		fmt.Printf("Error retrieving configuration file path for service '%s'\n", sn)
		return
	}
	fmt.Println(aOutput[1])

}

func init() {
	ServiceCmd.AddCommand(disableCmd)
	ServiceCmd.AddCommand(enableCmd)
	ServiceCmd.AddCommand(restartCmd)
	ServiceCmd.AddCommand(startCmd)
	ServiceCmd.AddCommand(stopCmd)
	ServiceCmd.AddCommand(maskCmd)
	ServiceCmd.AddCommand(unmaskCmd)
	//ServiceCmd.AddCommand(getconfCmd) TODO: Need to implement logic for getconf
	ServiceCmd.AddCommand(getservicefileCmd)
}
