package cmd

import (
	"fmt"
	"tadataka/encoder"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type settings struct {
	InputDir            string `json:"input_dir"`
	OutputDir           string `json:"output_dir"`
	LatColumn           int    `json:"lat_column"`
	LngColumn           int    `json:"lng_column"`
	GridColumnName      string `json:"grid_column_name"`
	NoHeader            bool   `json:"no_header"`
	OverwriteOutputfile bool   `json:"overwrite_outputfile"`
	Threads             int    `json:"threads"`
	SlackURL            string `json:"slackURL"`
}

func init() {
	olcCmd.PersistentFlags().String("config", "./config.json", "set config file path (JSON)")
	rootCmd.AddCommand(olcCmd)
}

var olcCmd = &cobra.Command{
	Use:   "olc",
	Short: "Devide geospatial CSV file with Open Location Code (OLC)",
	Run: func(cmd *cobra.Command, args []string) {

		color.Blue("TADATAKA OLC Encoder")
		configPath, err := cmd.PersistentFlags().GetString("config")
		if err != nil {
			fmt.Println("[TADATAKA] Flag Parse Error:", err)
			return
		}

		//bytes, err := ioutil.ReadFile(config)
		//TODO read JSON
		if err != nil {
			fmt.Println("[TADATAKA] JSON Reading Error:", err)
			return
		}

		//TODO implement single file mode and multiple file mode (directory mode)

		teststr := encoder.Encode(35.7720007, 139.7472105)
		fmt.Println(teststr)
		encoder.EncodeCSV(configPath)
	},
}
