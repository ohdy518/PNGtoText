package locator

import (
	"PNGtoText/imageHandler"
	"fmt"
	"strconv"
)

type RGBBin struct {
	rBin string
	gBin string
	bBin string
}

type pixelXY struct {
	x int64
	y int64
}

var initIndex int

func GetInitIndex() {
	getLocatorRGB()
}

func getLocatorRGB() {
	var err error

	var locator = imageHandler.RGBArray[0] // pixel 0, 0
	var locatorRGBBin RGBBin
	var locatorPixelXY pixelXY
	locatorRGBBin.rBin = fmt.Sprintf("%08b", locator.R)
	locatorRGBBin.gBin = fmt.Sprintf("%08b", locator.G)
	locatorRGBBin.bBin = fmt.Sprintf("%08b", locator.B)
	var concatBin string = locatorRGBBin.rBin + locatorRGBBin.gBin + locatorRGBBin.bBin
	var pixelXBin string = concatBin[:12]
	var pixelYBin string = concatBin[12:]
	locatorPixelXY.x, err = strconv.ParseInt(pixelXBin, 2, 0)
	if err != nil {
		panic(err)
	}
	locatorPixelXY.y, err = strconv.ParseInt(pixelYBin, 2, 0)
	if err != nil {
		panic(err)
	}

	fmt.Println(locatorPixelXY.x, locatorPixelXY.y)
}
