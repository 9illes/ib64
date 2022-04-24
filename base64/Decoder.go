package base64

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Decode decode base64 encoded string
func Decode(s string) (sDecoded string) {
	var indices []int = getIndices(s)
	indicesInbinaryNotation := indices2BinaryNotation(indices)
	stringBinary := strings.Join(indicesInbinaryNotation, "")
	bytes := extractBytes(stringBinary)
	sDecoded = binaryNotationToASCII(bytes)
	return
}

// getIndices return indice number of each char
func getIndices(s string) (indices []int) {
	var base64Index = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	for _, char := range s {
		if string(char) != "=" {
			indices = append(indices, strings.Index(base64Index, string(char)))
		}
	}
	return
}

// indices2BinaryNotation convert each indice to 6 bits binary notation
func indices2BinaryNotation(indices []int) (binaries []string) {
	for _, i := range indices {
		binaries = append(binaries, fmt.Sprintf("%06b", i))
	}
	return
}

// extractBytes split binary notation every 8 bits. Incomplete bytes are dropped.
func extractBytes(bs string) (bytes []string) {
	re := regexp.MustCompile(`.{8}`)
	bytes = append(bytes, re.FindAllString(bs, -1)...)
	return
}

// binaryNotationToASCII convert each byte (binary notation) in ASCII char
func binaryNotationToASCII(bytes []string) (sDecoded string) {
	for _, sByte := range bytes {
		i, _ := strconv.ParseInt(sByte, 2, 32)
		// fmt.Printf("%s > %d > %s\n", sByte, i, string(byte(i)))
		sDecoded += string(byte(i))
	}
	return
}
