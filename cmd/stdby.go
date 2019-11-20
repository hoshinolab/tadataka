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
		//TODO get version data from another file
		fmt.Println("LOAD DATA TO REDIS")
		db.RedisCSVLoader()
	},
}
