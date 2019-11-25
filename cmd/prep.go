package cmd

import (
	"tadataka/download"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(downloadCmd)
}

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Prepare for geocoding/reverse geocofing to download address data",
	Run: func(cmd *cobra.Command, args []string) {
		color.Blue("TADATAKA Preparing Tool")
		download.DownloadWizard()
	},
}
