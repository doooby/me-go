package cmd

// import (
// 	"fmt"
// 	"os"
// 	"me-go/internal/repository"
// 	"text/tabwriter"
// 
// 	"github.com/spf13/cobra"
// )
// 
// var all bool
// 
// var listCmd = &cobra.Command{
// 	Use:   "list",
// 	Short: "List tasks",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		tasks, err := repository.ListTasks(all)
// 		if err != nil {
// 			fmt.Println("Error fetching tasks:", err)
// 			return
// 		}
// 
// 		// Use tabwriter for pretty output
// 		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
// 		fmt.Fprintln(w, "ID\tCaption\tStarted\tStatus")
// 		fmt.Fprintln(w, "--\t-------\t-------\t------")
// 
// 		for _, t := range tasks {
// 			status := "Pending"
// 			if t.EndAt != nil {
// 				status = fmt.Sprintf("Done (%s)", t.EndAt.Format("15:04"))
// 			}
// 			fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", t.ID, t.Task, t.StartAt.Format("2006-01-02 15:04"), status)
// 		}
// 		w.Flush()
// 	},
// }
// 
// func init() {
// 	rootCmd.AddCommand(listCmd)
// 	listCmd.Flags().BoolVarP(&all, "all", "a", false, "Show all tasks including completed ones")
// }
