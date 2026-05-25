package data_structures

import (
	"Pokemon_Gen_2_Save_Reader/character_encoding"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

// Trainer struct stores the Name and TID of the Trainer.
type Trainer struct {
	// Name is the name of the Trainer
	Name [11]byte
	// TID is ID number of the Trainer
	TID [2]byte
}

// ReadTrainer reads the Trainer info from the save file.
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

// PrintTrainer prints the Name and TID of the given Trainer
func PrintTrainer(trainer *Trainer) {
	fmt.Println("Name:", character_encoding.ConvertString(trainer.Name[:]))
	fmt.Println("Trainer ID:", binary.BigEndian.Uint16(trainer.TID[:]))
}
