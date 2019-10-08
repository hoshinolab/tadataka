package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tadataka",
	Short: "Tadataka is a CLI tool designated to preprocess geo big data.",
	Long: `Tadataka is a CLI tool designated to preprocess geo big data.
				  developed by ryo-a`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TADATAKA")
	},
}

//exec
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
