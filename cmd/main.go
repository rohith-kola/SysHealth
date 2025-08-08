package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {

	fmt.Println("Start")

	// Main command
	rootCommand := &cobra.Command{
		Use:   "sysh",
		Short: "A simple CLI system health monitor application",
		Long: `A CLI application used to track the system health build using Golang cobra package.
		The output can also be exported.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello! Welcome to SysHealth!", "Use --help to see available commands")
		},
	}

	// Subcommand Version
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number",
		Long:  "Print the version number of this application",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("syshealth version 1.0.0")
		},
	}

	// Subcommand cpu
	var cpuCmd = &cobra.Command{
		Use:   "cpu",
		Short: "Show CPU information and usage",
		Long:  "Display CPU information including usage percentage, load averages, and core details",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
}
