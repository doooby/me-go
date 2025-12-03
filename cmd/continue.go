package cmd

import (
	"log"
	app "me-go/internal"
	"me-go/internal/model"
	"me-go/internal/repository"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	var cmd = &cobra.Command{
		Use:     "continue [id]",
		Aliases: []string{"c"},
		Short:   "Continue on a task",
		Args:    cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var id int64
			var task model.Task
			var err error

			if len(args) > 0 {
				id, err = app.StrToInt64(args[0])
				if err != nil {
					log.Fatalf("Error: %v", err)
				}
				task, err = repository.FindTaskById(id)
				if err != nil {
					log.Fatalf("Error: %v", err)
				}
			} else {
				task, err = repository.GetLastTask()
				if err != nil {
					log.Fatalf("Error: %v", err)
				}
			}

			repository.CreateTask(task.Task, task.Message, time.Now())
		},
	}

	rootCmd.AddCommand(cmd)
}
