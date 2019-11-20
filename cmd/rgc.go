package cmd

import (
	"fmt"
	"tadataka/rgc"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(rgcCmd)
}

var rgcCmd = &cobra.Command{
	Use:   "rgc",
	Short: "reverse geocoder",
	Run: func(cmd *cobra.Command, args []string) {
		color.Blue("TADATAKA Reverse Geocoder")
		inputFilePath := args[0]
		outputDirPath := args[1]

		//rgc.reverseGeocoder(inputFilePath, outputDirPath)
		fmt.Println(inputFilePath, outputDirPath)
		rgc.ReverseGeocodeCSV(inputFilePath, outputDirPath, 4, 3)

	},
}
