package util

import (
	"fmt"

	olc "github.com/google/open-location-code/go"
)

type Grids struct {
	East       string
	West       string
	South      string
	North      string
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

func GetNeighborGrid(grid string) (g Grids) {
	///2,3,4,5,6,7,8,9,C,F,G,H,J,M,P,Q,R,V,W,X
	//Validation
	if err := olc.Check(grid); err != nil {
		fmt.Println(err)
		return
	}

	//x axis and y axis
	x := "2"
	y := "2"

	//TODO Detection Level
	//TODO 端部だったら上位グリッドも取得
	return
}
