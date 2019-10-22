package encoder

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"tadataka/util"

	"github.com/fatih/color"
	olc "github.com/google/open-location-code/go"
)

func EncodeGridLevel(lat, lng float64, level int) string {
	olcGridName := olc.Encode(lat, lng, 6)[:6]
	return olcGridName
}

func EncodeCSV(path string) {
	color.Blue("OLC READER")
	st := util.JsonReader(path)
	fmt.Println(st.OutputDir)
}

//グリッド別のディレクトリを作る設計に変える
func EncodeSingleCSV(inputFilePath, outputDirPath string, latCol, lngCol int, header bool) {
	inputFile, inputErr := os.Open(inputFilePath)
	if inputErr != nil {
		panic(inputErr)
	}
	defer inputFile.Close()

	r := regexp.MustCompile(`.csv$`)
	fn := filepath.Base(r.ReplaceAllString(inputFilePath, "")) //file name without extension (.csv)
	outputFullPath := filepath.Join(outputDirPath, fn)
	if err := os.MkdirAll(outputFullPath, 0777); err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		p := util.CSVRowParser(line, 2, 3)
		grid := EncodeGridLevel(p.Lat, p.Lng, 6)
		gridCSVpath := filepath.Join(outputFullPath, grid+".csv")
		util.CSVRowWriter(line, gridCSVpath)
	}

}
