package main

import (
	"Pokemon_Gen_2_Save_Reader/data_structures"
	"log"
	"os"
)

func main() {
	fp, err := os.Open("debug/test.sav")
	if err != nil {
		log.Fatal(err)
	}

	trainer, err := data_structures.ReadTrainer(fp)
	if err != nil {
		println(err.Error())
	}
	data_structures.PrintTrainer(trainer)

	defer func(fp *os.File) {
		err := fp.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(fp)
}
