package common

import (
	"fmt"
	"image/color"
	"strconv"
	"strings"
)

type RGBBin struct {
	RBin string
	GBin string
	BBin string
}

type PixelXY struct {
	X int64
	Y int64
}

func TextToBinaryString(text string) string {
	var binaryStr string
	for _, c := range text {
		binaryStr += fmt.Sprintf("%08b", c)
	}
	return binaryStr
}

func BinaryStringToText(binaryStr string) (string, error) {
	var text string
	for i := 0; i < len(binaryStr); i += 8 {
		byteStr := binaryStr[i : i+8]
		charCode, err := strconv.ParseInt(byteStr, 2, 64)
		if err != nil {
			return "", err
		}
		text += string(rune(charCode))
	}
	return text, nil
}

func ConcatenateBinaryString(segments []string) string {
	return strings.Join(segments, "")
}

func ConcatenateFromRGB(rgba color.RGBA) string {
	var locator = rgba // pixel 1, 1
	var locatorRGBBin RGBBin
	locatorRGBBin.RBin = fmt.Sprintf("%08b", locator.R)
	locatorRGBBin.GBin = fmt.Sprintf("%08b", locator.G)
	locatorRGBBin.BBin = fmt.Sprintf("%08b", locator.B)
	return locatorRGBBin.RBin + locatorRGBBin.GBin + locatorRGBBin.BBin
}
