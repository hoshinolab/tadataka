package cmd

import (
	"fmt"
	"tadataka/db"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(stdbyCmd)
}

var stdbyCmd = &cobra.Command{
	Use:   "stdby",
	Short: "Load data to Redis",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Data Loader to Redis")
		db.RedisCSVLoader()
	},
}
