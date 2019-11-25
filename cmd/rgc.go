package cmd

import (
	"fmt"
	"tadataka/rgc"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	rgcCmd.PersistentFlags().Int("lat", 1, "Column number of latitude in CSV file. (begin from 0)")
	rgcCmd.PersistentFlags().Int("lng", 2, "Column number of longitude in CSV file. (begin from 0)")
	rootCmd.AddCommand(rgcCmd)
}

var rgcCmd = &cobra.Command{
	Use:   "rgc",
	Short: "Reverse geocoder",
	Run: func(cmd *cobra.Command, args []string) {
		color.Blue("Reverse Geocoder")
		inputFilePath := args[0]
		outputDirPath := args[1]

		latCol, err := cmd.PersistentFlags().GetInt("lat")
		if err != nil {
			fmt.Println("[TADATAKA] Flag Parse Error:", err)
			return
		}

		lngCol, err := cmd.PersistentFlags().GetInt("lng")
		if err != nil {
			fmt.Println("[TADATAKA] Flag Parse Error:", err)
			return
		}
		fmt.Println(inputFilePath, outputDirPath)
		rgc.ReverseGeocodeCSV(inputFilePath, outputDirPath, latCol, lngCol)

	},
}
