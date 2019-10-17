package encoder

import (
	"github.com/fatih/color"
	"github.com/google/open-location-code/go"
)

func Encode(lat, lng float64) string {
	olcGridName := olc.Encode(lat, lng, 10) //緯度(float64),経度(float64),OLCの桁数(int)
	return olcGridName
}

func EncodeCSV(path string) {
	color.Blue("OLC READER")
}
