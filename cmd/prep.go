package cmd

import (
	"tadataka/prep"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(prepCmd)
}

var prepCmd = &cobra.Command{
	Use:   "prep",
	Short: "Prepare for geocoding/reverse geocofing to download address data",
	Run: func(cmd *cobra.Command, args []string) {
		color.Blue("TADATAKA Preparing Tool")
		prep.DownloadWizard()
	},
}
