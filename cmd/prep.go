package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(prepCmd)
}

var prepCmd = &cobra.Command{
	Use:   "prep",
	Short: "Prepare for geocoding/reverse geocofing to download address data",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TADATAKA PREP")

		//TODO implement GSI file downloader
	},
}
