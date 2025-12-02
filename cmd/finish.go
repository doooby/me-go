package cmd

import (
	"github.com/spf13/cobra"
	"log"
	app "me-go/internal"
	"me-go/internal/model"
	"me-go/internal/repository"
	"time"
)

func init() {
	var endTimeArg string
	var cmd = &cobra.Command{
		Use:     "finish [id] [-e end_at]",
		Aliases: []string{"f"},
		Short:   "Finish a task",
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
				id, err = repository.FindUnfinishedId()
				if err != nil {
					log.Fatalf("Error: %v", err)
				}
				task, err = repository.FindTaskById(id)
				if err != nil {
					log.Fatalf("Error: %v", err)
				}
			}

      endAt := time.Now()
			if len(endTimeArg) > 0 {
				endAt, err = app.ParseShorthandTime(endTimeArg, endAt)
				if err != nil {
					log.Fatalf("Error: %v", err)
				}
			}

			repository.UpdateTaskEndAt(task.ID, endAt)
		},
	}

	cmd.Flags().StringVarP(&endTimeArg, "end-at", "e", "", "end time")
	rootCmd.AddCommand(cmd)
}
