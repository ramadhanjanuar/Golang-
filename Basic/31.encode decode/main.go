package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var data string = "john wick"

	var encodeString = base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("Encode: ", encodeString)

	var decodeByte, _ = base64.StdEncoding.DecodeString(encodeString)
	var decodeString = string(decodeByte)
	fmt.Println("Decode: ", decodeString)

	var encoded = make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(encoded, []byte(data))
	var encodedString = string(encoded)
	fmt.Println(encodedString)

	var decoded = make([]byte, base64.StdEncoding.DecodedLen(len(encoded)))
	var _, err = base64.StdEncoding.Decode(decoded, encoded)
	if err != nil {
		fmt.Println(err.Error())
	}

	var decodedString = string(decoded)
	fmt.Println(decodedString)

	var url = "hhtps://www.ramadhanjanuar.co.id"

	var encodedURL = base64.URLEncoding.EncodeToString([]byte(url))
	fmt.Println(encodedURL)

	var decodedURLByte, _ = base64.URLEncoding.DecodeString(encodedURL)
	var decodedURLString = string(decodedURLByte)
	fmt.Println(decodedURLString)

}
