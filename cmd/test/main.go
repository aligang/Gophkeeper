package main

import (
	"os"
)

func main() {
	f, err := os.OpenFile("/tmp/file132", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err.Error())
	}

	byteString := []byte("abcbasdf")
	for i := 0; i < len(byteString); i++ {
		f.Write([]byte{byteString[i]})
	}

}
