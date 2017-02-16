package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
)

func main() {
	bytes1 := make([]byte, 2)
	bytes1[0] = 0
	bytes1[1] = 1
	bytes2 := make([]byte, 3)
	bytes2[0] = 2
	bytes2[1] = 3
	bytes2[2] = 4
	// bytes3 := make([]byte, 3)
	// bytes3 = {3, 4, 5}
	superBytes := make([][]byte, 2)

	fmt.Println(superBytes)
	superBytes[0] = bytes1
	superBytes[1] = bytes2
	fmt.Println(superBytes)
	byteDelim := make([]byte, 1)
	byteDelim[0] = 0xff
	joined := bytes.Join(superBytes, byteDelim)
	fmt.Println(joined)
	joinedAndEncoded := hex.EncodeToString(joined)
	fmt.Println(joinedAndEncoded)
	joinedAndDecoded, err := hex.DecodeString(joinedAndEncoded)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(joinedAndDecoded)

	joinedAndSplit := bytes.Split(joinedAndDecoded, byteDelim)
	fmt.Println(joinedAndSplit)

	// fmt.Println(bytes.Split(bytes.Join(superBytes, byteDelim)), )
}
