/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package system

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	name string
	path string
)

// SystemCmd represents the system command
var SystemCmd = &cobra.Command{
	Use:   "system",
	Short: "System contains list of terminal commands",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Error: No arguments provided.")
		cmd.Help() // Prints the help message for the current command
		os.Exit(0)
	},
}

func finddir(name, path string) {
	cmd := exec.Command("find", path, "-type", "d", "-iname", "*"+name+"*")
	output, _ := cmd.CombinedOutput()
	fmt.Println(string(output))
}

func findfile(name, path string) {
	cmd := exec.Command("find", path, "-type", "f", "-iname", "*"+name+"*")
	output, _ := cmd.CombinedOutput()
	fmt.Println(string(output))
}

func init() {
	SystemCmd.AddCommand(finddirCmd)
	SystemCmd.AddCommand(findfileCmd)
}
