package prep

import (
	"os"
	"bufio"
	"fmt"
)

func DownloadJukyoJusho(){
	f, err := os.Open("./res/license_of_gsi_data.txt")
    if err != nil {
        panic(err)
    }
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
    f.Close()
}
