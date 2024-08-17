package locator

import (
	"PNGtoText/imageHandler"
	"fmt"
	"image/color"
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

var initIndex int64
var terminalIndex int64

var ITSlice []color.RGBA

func GetInitTerminalIndex() {
	initIndex = getIndex(0)
	terminalIndex = getIndex(1)

	// Debug override;
	// Comment the code below to use function normally.

	initIndex = 20
	terminalIndex = 24

	// Do not comment the code below.

	fmt.Println(initIndex)
	fmt.Println(terminalIndex)

}

func SetITSliceFromITIndex() {
	ITSlice = imageHandler.RGBArray[initIndex:terminalIndex]
	fmt.Println(ITSlice)
}

func getIndex(index int) int64 {
	var err error

	var locator = imageHandler.RGBArray[index] // pixel 1, 1
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

	return int64(imageHandler.Dimension.Width)*(locatorPixelXY.y) + locatorPixelXY.x
}
