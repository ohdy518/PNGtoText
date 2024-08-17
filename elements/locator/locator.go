package locator

import (
	"PNGtoText/common"
	"PNGtoText/imageHandler"
	"fmt"
	"image/color"
	"strconv"
)

var initIndex int64
var terminalIndex int64

var ITSlice []color.RGBA

func GetInitTerminalIndex() {
	initIndex = getIndex(0)
	terminalIndex = getIndex(1)

	// Debug override;
	// Comment the code below to use function normally.

	initIndex = 20
	terminalIndex = 22

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
	var locatorRGBBin common.RGBBin
	var locatorPixelXY common.PixelXY
	locatorRGBBin.RBin = fmt.Sprintf("%08b", locator.R)
	locatorRGBBin.GBin = fmt.Sprintf("%08b", locator.G)
	locatorRGBBin.BBin = fmt.Sprintf("%08b", locator.B)
	var concatBin string = locatorRGBBin.RBin + locatorRGBBin.GBin + locatorRGBBin.BBin
	var pixelXBin string = concatBin[:12]
	var pixelYBin string = concatBin[12:]
	locatorPixelXY.X, err = strconv.ParseInt(pixelXBin, 2, 0)
	if err != nil {
		panic(err)
	}
	locatorPixelXY.Y, err = strconv.ParseInt(pixelYBin, 2, 0)
	if err != nil {
		panic(err)
	}

	return int64(imageHandler.Dimension.Width)*(locatorPixelXY.Y) + locatorPixelXY.X
}
