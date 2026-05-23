package data_structures

import (
	"Pokemon_Gen_2_Save_Reader/character_encoding"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

type Trainer struct {
	Name [11]byte
	TID  [2]byte
}

func ReadTrainer(fp *os.File) (*Trainer, error) {
	var trainer Trainer

	// go to 0x2009 and read the TID
	_, err := fp.Seek(0x2009, io.SeekStart)
	if err != nil {
		return nil, err
	}

	tidData := make([]byte, 2)
	_, err = fp.Read(tidData)
	trainer.TID = [2]byte(tidData)

	// go to 0x200B and read the Trainer Name
	_, err = fp.Seek(0x200B, io.SeekStart)
	if err != nil {
		return nil, err
	}

	nameData := make([]byte, 11)
	_, err = fp.Read(nameData)
	trainer.Name = [11]byte(nameData)

	return &trainer, nil
}

func PrintTrainer(trainer *Trainer) {
	fmt.Println("Name:", character_encoding.ConvertString(trainer.Name[:]))
	fmt.Println("Trainer ID:", binary.BigEndian.Uint16(trainer.TID[:]))
}
