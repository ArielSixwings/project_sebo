package main

import (
	"gocv.io/x/gocv"
    //"path/filepath"
	"fmt"
	//"os/exec"
	"./imageprocessing"
)

func main() {

	

	Book := gocv.NewMat()     //mat for paint
	secondBook := gocv.NewMat()
	
	Book = imageprocessing.ReadImage(Book, "./imageprocessing/images/prt1.png", true, true, true)
	secondBook = imageprocessing.ReadImage(secondBook, "./imageprocessing/images/prt1.png", true, true, true)

	for r := 0; r < Book.rows(); r++ {
		for c := 0; c < _; c++ {
			sum += (Book[r][c] - secondBook[r][c])
		}
	}

	fmt.Printf("fim")
}