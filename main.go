package main

import (
	"PNGtoText/elements/locator"
	"PNGtoText/imageHandler"
	"os"
)

func main() {
	// Open the PNG file
	file, err := os.Open("./test/redblueflag.png")
	if err != nil {
		panic(err)
		return
	}
	defer func() {
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}()

	imageHandler.SetRGBArrayFromFile(file)

	locator.GetInitTerminalIndex()
	locator.SetITSliceFromITIndex()
}
