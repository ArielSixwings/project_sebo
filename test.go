package main

import (
	"fmt"
	"os/exec"
)

func main() {

	Book := gocv.NewMat()     //mat for paint

	Book = imageprocessing.ReadImage(Book, "./imageprocessing/images/sample.jpeg", true, true, true)


}