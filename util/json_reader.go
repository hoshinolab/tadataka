package libs

import (
	"encoding/json"
	"io/ioutil"

	"github.com/fatih/color"
)

type Settings struct {
	InputDir  string `json:"input_dir"`
	OutputDir string `json:"output_dir"`
	LatColumn int    `json:"lat_column"`
	LngColumn int    `json:"lng_column"`
	Header    bool   `json:"header"`
}

func JsonReader(jsonPath string) *Settings {
	bytes, err := ioutil.ReadFile(jsonPath)

	if err != nil {
		color.Red("TADATAKA JSON Reading Error:")
		panic(err)
	}

	jsonBytes := ([]byte)(bytes)
	settingData := new(Settings)

	if err := json.Unmarshal(jsonBytes, settingData); err != nil {
		color.Red("TADATAKA JSON Unmarshall Error:")
		panic(err)
	}
	return settingData
}
