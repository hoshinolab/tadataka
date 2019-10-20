package encoder

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"unsafe"

	"github.com/fatih/color"
	"github.com/google/open-location-code/go"
)

type settings struct {
	InputDir  string `json:"input_dir"`
	OutputDir string `json:"output_dir"`
	LatColumn int    `json:"lat_column"`
	LngColumn int    `json:"lng_column"`
	Header    bool   `json:"header"`
}

//TODO implement JSON Reader

func jsonReader(jsonPath string) *settings {
	bytes, err := ioutil.ReadFile(jsonPath)

	if err != nil {
		color.Red("TADATAKA JSON Reading Error:")
		panic(err)
	}

	jsonBytes := ([]byte)(bytes)
	settingData := new(settings)

	if err := json.Unmarshal(jsonBytes, settingData); err != nil {
		color.Red("TADATAKA JSON Unmarshall Error:")
		panic(err)
	}
	return settingData
}

func EncodeGridLevel(lat, lng float64) string {
	olcGridName := olc.Encode(lat, lng, 6)[:6]
	return olcGridName
}

func EncodeCSV(path string) {
	color.Blue("OLC READER")
	st := jsonReader(path)
	fmt.Println(st.OutputDir)
}

// ファイル内に記述する設計だが、グリッド別のディレクトリを作る設計に変える
// add Grid to CSV : add grid to each **SINGLE** CSV file
func addGridToCSVFile(inputFilePath string, outputFilePath string, latColumn int, lngColumn int, gridColumnName string, noHeader bool) {
	fmt.Println("Adding Grid: " + inputFilePath)

	inputFile, inputErr := os.Open(inputFilePath)
	if inputErr != nil {
		panic(inputErr)
	}
	defer inputFile.Close()

	outfile, outerr := os.OpenFile(outputFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if outerr != nil {
		log.Fatal(outerr)
	}
	defer outfile.Close()

	scanner := bufio.NewScanner(inputFile)
	headerLine := ""

	counter := 0
	w := bufio.NewWriter(outfile)
	for scanner.Scan() {
		line := scanner.Text()
		//ヘッダありの場合 ((TODO prelude で noheader だったものが header になって反転した
		if !noHeader && counter == 0 {
			headerLine = line
			fmt.Fprintln(outfile, headerLine+","+gridColumnName)
		} else {
			lineArray := strings.Split(line, ",")
			lineLat, _ := strconv.ParseFloat(lineArray[latColumn], 64)
			lineLng, _ := strconv.ParseFloat(lineArray[lngColumn], 64)
			grid := olc.Encode(lineLat, lineLng, 6)[:6]
			//fmt.Println(outputFilePath, lineLat, lineLng, grid)

			// ファイル内に記述する設計だが、グリッド別のディレクトリを作る設計に変える
			var outputStr = make([]byte, 0, 20)
			outputStr = append(outputStr, line...)
			outputStr = append(outputStr, ","...)
			outputStr = append(outputStr, grid...)
			outputStr = append(outputStr, "\n"...)
			w.WriteString(*(*string)(unsafe.Pointer(&outputStr))) //stringify
		}
		counter++
	}
	w.Flush()
	fmt.Println("Done: " + inputFilePath)
}
