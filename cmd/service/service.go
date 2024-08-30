/*
Copyright © 2024 AVINASH GHADSHI <avinashghadshi.official@gmail.com>
*/
package service

import (
	"bufio"
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
	fmt.Printf("Service file for service '%s':\n", sn)
	fmt.Println("--------------------------------")
	fmt.Println(aOutput[1])
	fmt.Println("--------------------------------")
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func readFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var content string
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}
	return content, nil
}

func getOS() string {
	if fileExists("/etc/lsb-release") {
		content, err := readFile("/etc/lsb-release")
		if err == nil && strings.Contains(content, "Ubuntu") {
			return "Ubuntu"
		}
	}

	if fileExists("/etc/redhat-release") {
		content, err := readFile("/etc/redhat-release")
		if err == nil && strings.Contains(content, "CentOS") {
			return "CentOS"
		}
	}

	return "Unknown"
}

func getLibraries(sn string) ([]byte, error) {
	os := getOS()
	switch os {
	case "Ubuntu":
		cmd := exec.Command("dpkg", "-L", sn)
		return cmd.CombinedOutput()

	case "CentOS":
		cmd := exec.Command("rpm", "-qc", sn)
		return cmd.CombinedOutput()

	default:
		return nil, fmt.Errorf("unsupported OS: %s", os)
	}
}

func init() {
	ServiceCmd.AddCommand(disableCmd)
	ServiceCmd.AddCommand(enableCmd)
	ServiceCmd.AddCommand(restartCmd)
	ServiceCmd.AddCommand(startCmd)
	ServiceCmd.AddCommand(stopCmd)
	ServiceCmd.AddCommand(maskCmd)
	ServiceCmd.AddCommand(unmaskCmd)
	ServiceCmd.AddCommand(getconfCmd)
	ServiceCmd.AddCommand(getservicefileCmd)
}
