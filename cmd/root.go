package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tadataka",
	Short: "tadataka is a CLI tool designated to preprocess geospatial big data.",
	Long: `tadataka is a CLI tool designated to preprocess geospatial big data.
developed by ryo-a (Keio Univ. Econ. Hoshino Lab. / RIKEN AIP Center)`,
	Run: func(cmd *cobra.Command, args []string) {
		color.Blue("tadataka: Geospatial Big Data Preprocessing Tool")
		fmt.Println("Please designate subcommand to run tadataka")
	},
}

//exec
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
