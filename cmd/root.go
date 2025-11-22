package cmd

import (
	"fmt"
	"os"
	"me-go/db"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task-cli",
	Short: "A simple CLI task manager",
}

// Execute is the entry point for the CLI
func Execute() {
	// Initialize DB file in the current directory
	db.InitDB("./var/db.sqlite")
	
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
