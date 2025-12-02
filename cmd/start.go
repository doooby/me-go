package cmd

import (
	"database/sql"
	"github.com/spf13/cobra"
	"log"
	"me-go/internal/repository"
	"time"
)

func init() {
	var cmd = &cobra.Command{
		Use:     "start task [message]",
		Aliases: []string{"s"},
		Short:   "Start a new task",
		Args:    cobra.RangeArgs(1, 2),
		Run: func(cmd *cobra.Command, args []string) {
			taskName := args[0]
			var message sql.NullString
			if len(args) > 1 {
				message.String = args[1]
				message.Valid = true
			}
			_, err := repository.CreateTask(taskName, message, time.Now())
			if err != nil {
				log.Fatalf("Error: %v", err)
			}
		},
	}

	rootCmd.AddCommand(cmd)
}
