package data_structures

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

// PokemonSpecies stores all the Species' Info like Pokedex Number, Name, Types and Base Stats.
type PokemonSpecies struct {
	// DexNum is the Pokedex number of the Species
	DexNum int
	// Name is the name of the Species
	Name string
	// Type1 is the first type of the Species
	Type1 string
	// Type2 is the second type of the Species
	Type2 string
	// BaseHP is the Base HP of the Species
	BaseHP int
	// BaseAtk is the Base Attack of the Species
	BaseAtk int
	// BaseDef is the Base Defense of the Species
	BaseDef int
	// BaseSpd is the Base Speed of the Species
	BaseSpd int
	// BaseSAtk is the Base Sp. Attack of the Species
	BaseSAtk int
	// BaseSDef is the Base Sp. Defense of the Species
	BaseSDef int
}

// FetchSpeciesInfo fetches the PokemonSpecies of a given species index
func FetchSpeciesInfo(index byte) *PokemonSpecies {
	file, _ := os.Open("pkmn2.csv")
	var speciesInfo PokemonSpecies
	reader := csv.NewReader(file)
	row := 0

	for {
		data, err := reader.Read()
		row++

		// terminate if EOF
		if err != nil {
			break
		}

		// skip headers
		if row == 1 {
			continue
		}

		// read species index
		readIndex, _ := strconv.ParseUint(data[0], 0, 8)

		// if the index is found, get all of its info
		if index == byte(readIndex) {
			speciesInfo.DexNum, _ = strconv.Atoi(data[1])
			speciesInfo.Name = data[2]
			speciesInfo.Type1 = data[3]
			speciesInfo.Type2 = data[4]
			speciesInfo.BaseHP, _ = strconv.Atoi(data[5])
			speciesInfo.BaseAtk, _ = strconv.Atoi(data[6])
			speciesInfo.BaseDef, _ = strconv.Atoi(data[7])
			speciesInfo.BaseSpd, _ = strconv.Atoi(data[8])
			speciesInfo.BaseSAtk, _ = strconv.Atoi(data[9])
			speciesInfo.BaseSDef, _ = strconv.Atoi(data[10])
		}
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	return &speciesInfo
}
