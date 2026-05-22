package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fp, err := os.Open("debug/test.sav")
	if err != nil {
		log.Fatal(err)
	}

	_, err = fp.Seek(0x2009, io.SeekStart)
	if err != nil {
		return
	}

	data := make([]byte, 2)
	TID, err := fp.Read(data)
	fmt.Println("Trainer ID:", binary.BigEndian.Uint16(data[:TID]))

	defer func(fp *os.File) {
		err := fp.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(fp)
}
