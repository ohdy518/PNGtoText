package common

import (
	"fmt"
	"strconv"
)

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
