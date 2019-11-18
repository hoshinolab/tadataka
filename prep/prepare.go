package prep

import (
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"

	//blank import for statik
	"tadataka/encoder"
	_ "tadataka/statik"
	"tadataka/util"

	"github.com/PuerkitoBio/goquery"
	"github.com/rakyll/statik/fs"
)

var jukyoJushoDir string

func DownloadWizard() {
	u, _ := user.Current()
	homeDir := u.HomeDir
	tadatakaDir := filepath.Join(homeDir, ".tadataka")
	jukyoJushoDir := filepath.Join(tadatakaDir, "jukyoJusho")
	jukyoJushoCSVPath := filepath.Join(jukyoJushoDir, "jukyo-jusho-concat.csv")

	statikFs, err := fs.New()
	if err != nil {
		panic(err)
	}

	//ISJ Data
	fISJ, err := statikFs.Open("/license_of_mlit_isj.txt")
	if err != nil {
		panic(err)
	}
	scannerISJ := bufio.NewScanner(fISJ)
	for scannerISJ.Scan() {
		fmt.Println(scannerISJ.Text())
	}
	fISJ.Close()

	if util.CLIQuestion() {
		fmt.Println("Download")
		// TODO ISJ download
		downloadIsj()
	} else {
		fmt.Println("Abort.")
	}

	//GSI Data
	f, err := statikFs.Open("/license_of_gsi_data.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	f.Close()

	if util.CLIQuestion() {
		fmt.Println("Download")
		downloadJukyoJusho()
		//TODO add grid to jukyojusho CSV
		encoder.AddGridToCSV(jukyoJushoCSVPath, jukyoJushoDir, 9, 8, false)
	} else {
		fmt.Println("Abort.")
	}
}

func downloadJukyoJusho() {
	u, _ := user.Current()
	homeDir := u.HomeDir
	tadatakaDir := homeDir + "/.tadataka"
	jukyoJushoDir = tadatakaDir + "/jukyoJusho"

	if f, err := os.Stat(tadatakaDir); os.IsNotExist(err) || !f.IsDir() {
		if err := os.MkdirAll(tadatakaDir, 0777); err != nil {
			panic(err)
		}
	}
	if f, err := os.Stat(jukyoJushoDir); os.IsNotExist(err) || !f.IsDir() {
		if err := os.MkdirAll(jukyoJushoDir, 0777); err != nil {
			panic(err)
		}
	}

	//TODO DOWNLOAD
	fmt.Println("Output destination directory: " + jukyoJushoDir)

	gsiDownloadURL := "https://saigai.gsi.go.jp/jusho/download/"

	doc, err := goquery.NewDocument(gsiDownloadURL)
	if err != nil {
		fmt.Print("Connection Error")
	}

	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		text := s.Text()
		if strings.Contains(url, "pref/") {
			getPref(gsiDownloadURL+"/"+url, text)
			time.Sleep(2000 * time.Millisecond)
		}
	})

	// unzip
	files, _ := ioutil.ReadDir(jukyoJushoDir)
	for _, file := range files {
		if strings.Contains(file.Name(), ".zip") {
			zp := path.Join(jukyoJushoDir, file.Name())
			extractCSV(zp)
		}
	}
	//create concat CSV
	ccf := "jukyo-jusho-concat.csv"
	ccp := path.Join(jukyoJushoDir, ccf)
	concatCSV, err := os.OpenFile(ccp, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer concatCSV.Close()

	// list up concat unzipped csv
	files, _ = ioutil.ReadDir(jukyoJushoDir)
	for _, file := range files {
		if strings.Contains(file.Name(), ".csv") && !strings.Contains(file.Name(), "concat") {
			srcbyte := []byte(file.Name())
			rx := regexp.MustCompile(".+_(.+)_(.+).csv")
			match := rx.FindSubmatch(srcbyte)
			prefName := string(match[1])
			cityName := string(match[2])

			//add csv data to concat CSV
			cp := path.Join(jukyoJushoDir, file.Name())
			cf, err := os.Open(cp)
			fmt.Println("open: " + cp)
			if err != nil {
				log.Fatal(err)
				return
			}

			s := bufio.NewScanner(transform.NewReader(cf, japanese.ShiftJIS.NewDecoder()))
			for s.Scan() {
				sl := strings.SplitN(s.Text(), ",", 2)
				writestr := sl[0] + "," + prefName + "," + cityName + "," + sl[1]
				fmt.Fprintln(concatCSV, writestr)
			}
			fmt.Println("close: " + cp)
		}
	}

	fmt.Println("DELETE TMP FILES")
	delFiles, _ := ioutil.ReadDir(jukyoJushoDir)
	for _, file := range delFiles {
		if strings.Contains(file.Name(), ".csv") && !strings.Contains(file.Name(), "concat") || strings.Contains(file.Name(), ".zip") {
			rp := path.Join(jukyoJushoDir, file.Name())
			if err := os.RemoveAll(rp); err != nil {
				fmt.Println(err)
			}
		}
	}

}

// get list of zip files from prefecture page
func getPref(url string, prefName string) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Print("Connection Error")
	}
	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		cityName := s.Text()
		if strings.Contains(url, ".zip") {
			fn := getZIPFileName(url)
			num := strings.Replace(fn, ".zip", "", -1)
			fu := getZIPFileURL(fn)
			fmt.Println(prefName + cityName + fu)
			downloadAndWait(num+"_"+prefName+"_"+cityName+".zip", fu)
		}
	})
}

// get a zip file name from path like "../data/03482.zip"
func getZIPFileName(path string) string {
	s := strings.Split(path, "/")
	return s[2]
}

// get a zip file URL from file name like "03482.zip"
func getZIPFileURL(filename string) string {
	return "https://saigai.gsi.go.jp/jusho/download/data/" + filename
}

func downloadAndWait(saveFileName, url string) error {
	res, err := http.Get(url)

	//error
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// create file for the download target
	out, err := os.Create(filepath.Join(jukyoJushoDir, saveFileName))
	if err != nil {
		return err
	}
	defer out.Close()

	// write response body to the file
	_, err = io.Copy(out, res.Body)

	fmt.Println("Done. Waiting")
	time.Sleep(2000 * time.Millisecond)
	return err
}

func extractCSV(src string) error {
	srcbyte := []byte(src)
	rx := regexp.MustCompile("\\d+_(.+)_(.+).zip")
	match := rx.FindSubmatch(srcbyte)
	prefName := string(match[1])
	cityName := string(match[2])

	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()
	for _, f := range r.File {
		fn := f.FileInfo().Name()
		if strings.Contains(fn, ".csv") {
			fmt.Println(fn)
			rf, err := f.Open()
			if err != nil {
				return err
			}
			defer rf.Close()

			buf := make([]byte, f.UncompressedSize)
			_, err = io.ReadFull(rf, buf)
			if err != nil {
				return err
			}

			savefn := strings.Replace(fn, ".csv", "_"+prefName+"_"+cityName+".csv", -1)
			path := filepath.Join(jukyoJushoDir, savefn)
			err = ioutil.WriteFile(path, buf, f.Mode())
			if err != nil {
				return err
			}

		}
	}
	return nil
}

func downloadIsj() {

}

func downloadIsjZip(url string) {
	url := "http://nlftp.mlit.go.jp/isj/dls/data/17.0a/" + "05000-17.0a.zip"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Referer", "http://nlftp.mlit.go.jp/cgi-bin/isj/dls/_download_files.cgi")

	client := new(http.Client)
	resp, _ := client.Do(req)

	byteArray, _ := ioutil.ReadAll(resp.Body)
	_, filename := path.Split(url)

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		file.Close()
	}()

	file.Write(byteArray)
}

func extractIsjCSV(path string) {

}

func concatIsjCSV(path string) {

}
