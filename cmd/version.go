package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Tadataka",
	Run: func(cmd *cobra.Command, args []string) {
		//TODO get version data from another file
		fmt.Println("TADATAKA Ver0.x")
	},
}
