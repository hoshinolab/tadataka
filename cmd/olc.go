package cmd

import (
	"fmt"
	"tadataka/encoder"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	olcCmd.PersistentFlags().String("config", "", "set config file path (JSON)")
	olcCmd.PersistentFlags().Int("lat", 1, "Column number of latitude in CSV file. (begin from 0)")
	olcCmd.PersistentFlags().Int("lng", 2, "Column number of longitude in CSV file. (begin from 0)")
	olcCmd.PersistentFlags().Int("buffer", 100000, "Buffering size (the number of rows)")
	olcCmd.PersistentFlags().Bool("header", true, "Whether CSV files have a header row or not. (default: true)")
	rootCmd.AddCommand(olcCmd)
}

var olcCmd = &cobra.Command{
	Use:   "olc",
	Short: "Subdivide geospatial CSV file with Open Location Code (OLC)",
	Run: func(cmd *cobra.Command, args []string) {

		color.Blue("TADATAKA OLC Encoder")
		configPath, err := cmd.PersistentFlags().GetString("config")
		if err != nil {
			fmt.Println("[TADATAKA] Flag Parse Error:", err)
			return
		}

		if configPath == "" {
			//TODO
			fmt.Println("Single File Mode")
			if len(args) != 2 {
				color.Red("[TADATAKA] Error! You must designate both input and output file path.")
				panic("terminated...")
			}
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

			bufferSize, err := cmd.PersistentFlags().GetInt("buffer")
			if err != nil {
				fmt.Println("[TADATAKA] Flag Parse Error:", err)
				return
			}

			header, err := cmd.PersistentFlags().GetBool("header")
			if err != nil {
				fmt.Println("[TADATAKA] Flag Parse Error:", err)
				return
			}

			encoder.EncodeSingleCSV(inputFilePath, outputDirPath, latCol, lngCol, header)

		} else {
			//TODO implement single file mode and multiple file mode (directory mode)
			fmt.Println("Multiple File Mode")
			encoder.EncodeCSV(configPath)
		}

	},
}
