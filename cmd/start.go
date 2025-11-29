package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"me-go/internal/repository"
)

func init() {
	var cmd = &cobra.Command{
		Use:   "start task [message]",
		Short: "Start a new task",
		Args:  cobra.RangeArgs(1, 2),
		Run: func(cmd *cobra.Command, args []string) {
			task := args[0]
			var message string
			if len(args) > 1 {
				message = args[1]
			}
			err := repository.CreateTask(task, message)
			if err != nil {
				log.Fatalf("Error: %v", err)
			}
		},
	}

	rootCmd.AddCommand(cmd)
}
