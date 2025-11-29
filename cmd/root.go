package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"me-go/db"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "me cmd ...",
	Short: "Meassures my time",
}

func Execute() {
	db.InitDB()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
