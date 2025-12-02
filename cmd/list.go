package cmd

import (
	"fmt"
	"log"
	"me-go/db"
	"me-go/internal/repository"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	var cmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List tasks (last 10)",
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			pagination := db.FirstTenPagination()
			tasks, err := repository.ListTasks(pagination)
			if err != nil {
				log.Fatalf("Error: %v", err)
			}

			colSizes := [4]int{}
			rows := [][4]string{}

			// gather the texts, meassure sizes
			for _, task := range tasks {
				duration, err := task.DurationText()
				if err != nil {
					log.Fatalf("Error: %v", err)
				}
				row := [4]string{
					strconv.FormatInt(task.ID, 10),
					task.Task,
					duration,
					task.Message.String,
				}
				rows = append(rows, row)
				for column, text := range row[:3] {
					colSizes[column] = max(colSizes[column], len(text))
				}
			}

			// fmt.Printf("rows:\n%#v\n", rows)

			// print
			for _, row := range rows {
				var paddedTexts [4]string
				for column := range row {
					paddedTexts[column] = fmt.Sprintf("%-*s", colSizes[column], row[column])
				}
				fmt.Println(strings.Join(paddedTexts[:], " | "))
			}
		},
	}

	rootCmd.AddCommand(cmd)
}
