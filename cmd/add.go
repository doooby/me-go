package cmd

import (
	"fmt"
	"log"
	"me-go/internal/repository"

	"github.com/spf13/cobra"
)

var caption string
var text string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Run: func(cmd *cobra.Command, args []string) {
		if caption == "" {
			fmt.Println("Error: Caption is required")
			return
		}
		err := repository.CreateTask(caption, text)
		if err != nil {
			log.Fatalf("Error creating task: %v", err)
		}
		fmt.Println("Task added successfully!")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&caption, "caption", "c", "", "Caption of the task (required)")
	addCmd.Flags().StringVarP(&text, "text", "t", "", "Detailed text of the task")
}
