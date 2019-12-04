package util

import (
	"fmt"

	olc "github.com/google/open-location-code/go"
)

type Grids struct {
	East       string
	NorthEast  string
	North      string
	NorthWest  string
	West       string
	SouthWest  string
	South      string
	SouthEast  string
	UpperLevel string
}

var incrementNeighborCode map[string]string = map[string]string{
	"2": "3",
	"3": "4",
	"4": "5",
	"5": "6",
	"6": "7",
	"7": "8",
	"8": "9",
	"9": "C",
	"C": "F",
	"F": "G",
	"G": "H",
	"H": "J",
	"J": "M",
	"M": "P",
	"P": "Q",
	"Q": "R",
	"R": "V",
	"V": "W",
	"W": "X",
	"X": "2",
}

var decrementNeighborCode map[string]string = map[string]string{
	"2": "X",
	"3": "2",
	"4": "3",
	"5": "4",
	"6": "5",
	"7": "6",
	"8": "7",
	"9": "8",
	"C": "9",
	"F": "C",
	"G": "F",
	"H": "G",
	"J": "H",
	"M": "J",
	"P": "M",
	"Q": "P",
	"R": "Q",
	"V": "R",
	"W": "V",
	"X": "W",
}

// only support "8Q7XPX2P+" level
func GetNeighborGrid(grid string) (g Grids) {
	///2,3,4,5,6,7,8,9,C,F,G,H,J,M,P,Q,R,V,W,X
	//Validation
	if err := olc.Check(grid); err != nil {
		fmt.Println(err)
		return
	}
	//8Q7XPX2P+J4 -> 8Q 7X PX 2P
	/*containsPlus := false
	if strings.Contains(grid, "+") {
		containsPlus = true
		grid = grid[:len(grid)-1] //remove "+"
	}*/

	// 1段階 TODO 書き直す
	// 8Q7XPX2P+J4 -> 8Q 7X PX 2Pの分割やって 4th 3rd 2nd 1st で下から判定してそれぞれ隣接取る
	// X, Y 別々に取るほうが良さそう

	firstLvY := grid[:1]
	firstLvX := grid[1:2]
	secondLvY := grid[2:3]
	secondLvX := grid[3:4]
	thirdLvY := grid[4:5]
	thirdLvX := grid[5:6]
	fourthLvY := grid[6:7]
	fourthLvX := grid[7:8]

	//4th
	//fourthNorthY := incrementNeighborCode[fourthLvY]
	//fourthSouthY := decrementNeighborCode[fourthLvY]
	if fourthLvY == "X" {
		// TODO check 3rd
		if thirdLvY == "X" {
			// TODO check 2nd
			if secondLvY == "X" {
				// TODO check 1st
				if firstLvY == "X" {
					// TODO check
				}
			}
		}
	}
	//TODO also chrck 2

	if fourthLvX == "X" || fourthLvX == "2" {
		// TODO check 3rd
		if thirdLvX == "X" || thirdLvX == "2" {
			// TODO check 2nd
			if secondLvX == "X" || secondLvX == "2" {
				// TODO check 1st
				if firstLvX == "X" || firstLvX == "2" {
					// TODO check
				}
			}
		}
	}

	//test[len(test)-2:len(test)]

	//TODO 端部だったら上位グリッドも取得
	return
}
