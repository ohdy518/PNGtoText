package parser

import (
	"PNGtoText/common"
	"fmt"
	"image/color"
)

func Parse(parseData []color.RGBA) {
	var binConcat string = concat(parseData)
	var result, err = common.BinaryStringToText(binConcat)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func concat(concatData []color.RGBA) string {
	var concatList []string = make([]string, 0)
	for _, i := range concatData {
		concatList = append(concatList, common.ConcatenateFromRGB(i))
	}
	return common.ConcatenateBinaryString(concatList)
}
