/*
Copyright © 2024 AVINASH GHADSHI <avinashghadshi.official@gmail.com>
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
	fmt.Println("Your Search Results: ")
	fmt.Println("--------------------------------")
	fmt.Println(string(output))
	fmt.Println("--------------------------------")
}

func findfile(name, path string) {
	cmd := exec.Command("find", path, "-type", "f", "-iname", "*"+name+"*")
	output, _ := cmd.CombinedOutput()
	fmt.Println("Your Search Results: ")
	fmt.Println("--------------------------------")
	fmt.Println(string(output))
	fmt.Println("--------------------------------")
}

func init() {
	SystemCmd.AddCommand(finddirCmd)
	SystemCmd.AddCommand(findfileCmd)
}
