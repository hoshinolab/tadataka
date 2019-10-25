package encoder

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"tadataka/util"

	"github.com/fatih/color"
	olc "github.com/google/open-location-code/go"
)

func EncodeGridLevel(lat, lng float64, level int) string {
	olcGridName := olc.Encode(lat, lng, level)[:level]
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

	buf := make(map[string]string, 150000)
	bufCount := 0

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		p := util.CSVRowParser(line, 2, 3)            //8Q7XJPXR+MM
		fullGrid := EncodeGridLevel(p.Lat, p.Lng, 11) //6桁で取ってるが、full gridが要るので6桁に絞らない　ファイル名用のgridはあとで[:6]で取る
		shortGrid := fullGrid[:6]
		//gridCSVpath := filepath.Join(outputFullPath, grid+".csv")
		var outputStr = make([]byte, 0, 30)
		outputStr = append(outputStr, line...)
		outputStr = append(outputStr, ","...)
		outputStr = append(outputStr, fullGrid...)
		outputStr = append(outputStr, "\n"...)
		outputLine := string(outputStr)

		bufArray := []string{buf[shortGrid], outputLine, "\r\n"}
		buf[shortGrid] = strings.Join(bufArray, "")
		bufCount++
		if bufCount > 1000 {
			fmt.Println("FLUSH") //TODO FLUSH with goroutine
		}
		//util.CSVRowWriter(line, gridCSVpath)
	}
	for keyGrid, csvData := range buf {
		fmt.Println(keyGrid)
		fmt.Println(csvData)
		fmt.Println("==============")
	}

}
