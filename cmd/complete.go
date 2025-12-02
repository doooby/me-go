package cmd

// import (
// 	"fmt"
// 	"strconv"
// 	"me-go/internal/repository"
// 
// 	"github.com/spf13/cobra"
// )
// 
// var completeCmd = &cobra.Command{
// 	Use:   "complete [id]",
// 	Short: "Mark a task as completed",
// 	Args:  cobra.ExactArgs(1), // Requires exactly one argument (the ID)
// 	Run: func(cmd *cobra.Command, args []string) {
// 		id, err := strconv.Atoi(args[0])
// 		if err != nil {
// 			fmt.Println("Invalid ID provided")
// 			return
// 		}
// 
// 		err = repository.CompleteTask(id)
// 		if err != nil {
// 			fmt.Println("Error completing task:", err)
// 			return
// 		}
// 		fmt.Printf("Task %d marked as complete.\n", id)
// 	},
// }
// 
// func init() {
// 	rootCmd.AddCommand(completeCmd)
// }
