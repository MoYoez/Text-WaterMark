package addWaterMark

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// how it works? https://medium.com/@umpox/be-careful-what-you-copy-invisibly-inserting-usernames-into-text-with-zero-width-characters-18b4e6f17b66

// textToBinary just type your name,as the watermark.
func textToBinary(username string) string {
	binaryStrings := make([]string, len(username))
	for i, char := range username {
		binaryStrings[i] = fmt.Sprintf("%08s", strconv.FormatInt(int64(int(char)), 2))
	}
	return strings.Join(binaryStrings, " ")
}

// binaryToZeroWidth Already Add WaterMark.
func binaryToZeroWidth(binary string) string {
	var zeroWidthBuilder strings.Builder
	for _, binaryNum := range binary {
		if binaryNum == '1' {
			zeroWidthBuilder.WriteRune('\u200B') // zero-width space
		} else if binaryNum == '0' {
			zeroWidthBuilder.WriteRune('\u200C') // zero-width non-joiner
		} else {
			zeroWidthBuilder.WriteRune('\u200D') // zero-width joiner
		}
	}
	return zeroWidthBuilder.String()
}

// zeroWidthToBinary decode to binary code.
func zeroWidthToBinary(encodingMsg string) string {
	var binaryBuilder string
	for _, char := range encodingMsg {
		switch char {
		case '\u200B': // zero-width space
			binaryBuilder += "1"
		case '\u200C': // zero-width non-joiner
			binaryBuilder += "0"
		default:
			binaryBuilder += " " // add single space
		}
	}
	return binaryBuilder
}

func DecodeBinaryWaterMark(str string) string {
	binaryArr := strings.Split(str, " ")
	textArr := []byte{}
	for _, binaryStr := range binaryArr {
		binary, err := strconv.ParseInt(binaryStr, 2, 64)
		if err != nil {
			panic(err)
		}
		textArr = append(textArr, byte(binary))
	}
	return string(textArr)
}

func AddWaterMarkToText(s string, watermark string) string {
	binaryCode := textToBinary(watermark)
	encodingName := binaryToZeroWidth(binaryCode)
	EncodingMsg := s[:] + encodingName
	return EncodingMsg
}

func getWaterMark(str string) string {
	usernamePattern := `\p{C}+`
	re := regexp.MustCompile(usernamePattern)
	waterMark := re.FindAllString(str, -1)
	BinaryWaterMark := zeroWidthToBinary(waterMark[0])
	binarytoText := DecodeBinaryWaterMark(BinaryWaterMark)
	return binarytoText
}
