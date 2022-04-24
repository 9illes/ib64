package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/9illes/ib64/base64"
)

func main() {

	var enc bool
	var dec bool

	flag.BoolVar(&enc, "e", false, "Encode string")
	flag.BoolVar(&dec, "d", false, "Decode string")

	flag.Parse()

	s := os.Args[len(os.Args)-1]

	var encoded string
	var decoded string

	if enc {
		encoded = base64.Encode(s)
		decoded = base64.Decode(encoded)
		fmt.Printf("base64.Encode(%s) => %s\nbase64.Decode(%s) => %s\n", s, encoded, encoded, decoded)
	} else if dec {
		decoded = base64.Decode(s)
		encoded = base64.Encode(decoded)
		fmt.Printf("base64.Decode(%s) => %s\nbase64.Encode(%s) => %s\n", s, decoded, decoded, encoded)
	} else {
		fmt.Println("Usage :\n  encoding : go run main.go -e foo\n  decoding : go run main.go -d Zm9vA")
	}

}
