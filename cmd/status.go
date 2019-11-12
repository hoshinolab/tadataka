package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(statusCmd)
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show status of TADATAKA",
	Run: func(cmd *cobra.Command, args []string) {
		//TODO show Redis Status, prep data status
		fmt.Println("LOAD DATA TO REDIS")
	},
}
