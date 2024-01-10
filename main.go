package main

import (
	"fmt"
	"os"
)

func main() {
	rdr, err := os.Open("nine.bd")
	if err != nil {
		panic(err)
	}
	board := NewBoardFromFile(9, rdr)
	for _, s := range board.Display() {
		fmt.Println(s)
	}
}
