package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(olcCmd)
}

var olcCmd = &cobra.Command{
	Use:   "olc",
	Short: "Open Location Code (OLC)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TADATAKA OLC Encoder")
	},
}
