package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ParsedCSV struct {
	Lat float64
	Lng float64
}

func CSVRowParser(csvRow string, latCol, lngCol int) *ParsedCSV {

	csvData := &ParsedCSV{Lat: 1, Lng: 2}

	slice := strings.Split(csvRow, ",")

	csvData.Lat, _ = strconv.ParseFloat(slice[latCol], 64)
	csvData.Lng, _ = strconv.ParseFloat(slice[lngCol], 64)

	return csvData
}

func CSVRowWriter(csvRow, outputCSVPath string) {
	file, err := os.OpenFile(outputCSVPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Fprintln(file, csvRow)
}

func CSVBufWriter(csvBuf, outputCSVPath string) {
	file, err := os.OpenFile(outputCSVPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(file)
	w := bufio.NewWriter(writer)
	w.WriteString(csvBuf)
	w.Flush()
	defer file.Close()
}
