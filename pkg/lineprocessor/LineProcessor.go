package lineprocessor

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

const HEADER = "X-ImunifyEmail-Filter-Info:"

type LineProcessor struct {
	readingEncodingMode bool
	b64builder          strings.Builder
}

func NewProcessor() LineProcessor {
	lp := LineProcessor{}
	lp.readingEncodingMode = false
	return lp
}

func (lp *LineProcessor) ProcessLine(input string) (string, bool) {
	if lp.readingEncodingMode {
		if checkIfBase64Continues(input) {
			lp.b64builder.WriteString(strings.TrimSpace(input))
			return "", false
		} else {
			lp.readingEncodingMode = false
			resStr := lp.b64builder.String()
			lp.b64builder.Reset()
			return HEADER + " " + decode(resStr), true
		}
	} else if strings.HasPrefix(input, HEADER) {
		//fmt.Println("found in line: ", str)
		lp.readingEncodingMode = true
		b64l := strings.TrimSpace(input[len(HEADER):])
		//fmt.Println("BASE64 Starts: ", b64l)
		lp.b64builder.WriteString(b64l)
		return "", false
	} else {
		return input, true
	}
}

func checkIfBase64Continues(str string) bool {
	return strings.HasPrefix(str, "\t")
}

func decode(b64str string) string {
	//log.Println("To Decode: ", b64str)
	dst, err := base64.StdEncoding.DecodeString(b64str)
	if err != nil {
		fmt.Println("Error decoding base64: ", err)
		os.Exit(1)
	}
	return string(dst)
}
