package encoder

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unsafe"

	"tadataka/util"

	"github.com/fatih/color"
	olc "github.com/google/open-location-code/go"
)

func EncodeGridLevel(lat, lng float64) string {
	olcGridName := olc.Encode(lat, lng, 6)[:6]
	return olcGridName
}

func EncodeCSV(path string) {
	color.Blue("OLC READER")
	st := util.JsonReader(path)
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
