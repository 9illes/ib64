package base64

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Encode string with Base64 algorithm
func Encode(s string) (sEncoded string) {

	var binaryNotation string = stringToBinaryNotation(s)
	bitsGroups := createStringsOf6bits(binaryNotation)
	byteGroups := padBits(bitsGroups)
	sEncoded = indices2string(byteGroups) + padding(len(s))
	return
}

func padding(length int) (padding string) {
	if length%3 == 0 {
		return
	}
	padding = strings.Repeat("=", 3-length%3)
	return
}

// stringToBinaryNotation convert each char to its binary notation (8 bits)
func stringToBinaryNotation(s string) (binaryString string) {
	for _, c := range s {
		binaryString += fmt.Sprintf("%08s", fmt.Sprintf("%b", c))
	}
	return
}

// createStringsOf6bits create a slice of strings of 6 bits (ex: 11000)
func createStringsOf6bits(bs string) (groups []string) {
	re := regexp.MustCompile(`.{6}`)
	bs += strings.Repeat("0", 6-len(bs)%6)
	groups = append(groups, re.FindAllString(bs, -1)...)
	return
}

// padBits pad each groups of 6 bits to get a byte
func padBits(groups []string) (bytesGroups []string) {
	for _, bits := range groups {
		bytesGroups = append(bytesGroups, fmt.Sprintf("00"+bits))
	}
	return
}

func indices2string(indices []string) (s string) {
	var base64Index = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	for _, byteStr := range indices {
		if i, err := strconv.ParseInt(byteStr, 2, 64); err != nil {
			fmt.Println(err)
		} else {
			s += string(base64Index[i])
		}
	}
	return
}
