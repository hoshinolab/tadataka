package rgc

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path"
	"strconv"
	"strings"
	"tadataka/db"
	"tadataka/encoder"
)

const (
	//Based on WGS-84 Earth ellipsoid
	R_x = 6378137.0           // Equatorial radius
	R_y = 6356752.3142        //Polar radius
	E   = 0.08181919092890633 //math.Sqrt( (R_x*R_x - R_y*R_y) / (R_x*R_x) )
	RAD = math.Pi / 180
)

func ReverseGeocodeCSV(inputCSVPath, outputDirPath string, latCol, lngCol int) {

	inputCSV, err := os.Open(inputCSVPath)
	if err != nil {
		panic(err)
	}

	//TODO make and open output CSV

	outputCSVFileName := "rgc.csv"
	outputCSVPath := path.Join(outputDirPath, outputCSVFileName)
	outputCSV, err := os.OpenFile(outputCSVPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	bs := bufio.NewScanner(inputCSV)
	for bs.Scan() {
		sl := strings.Split(bs.Text(), ",")
		lat, _ := strconv.ParseFloat(sl[latCol], 64)
		lng, _ := strconv.ParseFloat(sl[lngCol], 64)
		grid := encoder.EncodeGridLevel(lat, lng, 9)
		possibleAddress := db.GetMembersFromList(grid, "ISJ")
		/*
			format
			JukyoJusho
			岡山県:岡山市南区:泉田二丁目 : 1 : 34 : 34.627437313:133.920691406:8Q6MJWGC+X7
			ISJ
			岡山県:岡山市南区:泉田三丁目 :   : 2  : 34.625610:133.920917:8Q6MJWGC+69
		*/
		minDist := 99999.99 //memo
		minDistAddr := "NA"
		for _, dbVal := range possibleAddress {
			sval := strings.Split(dbVal, ":")
			possibleLat, _ := strconv.ParseFloat(sval[5], 64)
			possibleLng, _ := strconv.ParseFloat(sval[6], 64)
			possibleDist := GetDistance(lat, lng, possibleLat, possibleLng)
			if possibleDist < minDist {
				minDist = possibleDist
				minDistAddr = sval[0] + sval[1] + sval[2] + sval[3] + "-" + sval[4]
			}
		}
		fmt.Fprintln(outputCSV, bs.Text()+","+minDistAddr)

		//TODO Write to output CSV
	}
}

func GetDistance(originLat, originLng, destLat, destLng float64) float64 {
	// Hubeny Standard Formula
	// D = \sqrt{(D_y \times M)^2 + (D_x \times N \times \cos P)^2}
	// D_x : difference of longitude between two coordinates
	// D_y : difference of latitude between two coordinates
	// P: average of two latitudes
	// M = \frac{R_x(1-E^2)}{W^3}: Curvature radius of meridian
	// W = \sqrt{1-E^2 \times \sin^2 P}
	// N : Curvature radius of Prime vertical
	// E = \sqrt{\frac{R_x^2-R_y^2}{R_x^2}} : Eccentricity
	// R_x: Equatorial radius
	// R_y: Polar radius
	originLatRad := originLat * math.Pi / 180
	originLngRad := originLng * math.Pi / 180
	destLatRad := destLat * math.Pi / 180
	destLngRad := destLng * math.Pi / 180
	D_x := math.Abs(destLngRad - originLngRad)
	D_y := math.Abs(destLatRad - originLatRad)
	P := (destLatRad + originLatRad) / 2.0
	W := math.Sqrt(1 - E*E*math.Sin(P)*math.Sin(P))
	M := (R_x + (1 - E*E)) / W * W * W
	N := R_x / W
	D := math.Sqrt((D_y*M)*(D_y*M) + (D_x*N*math.Cos(P))*(D_x*N*math.Cos(P)))

	return D
}
