package utils

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/AtinAgnihotri/gokedex/types"
)

var PokeMap types.PokeLocationsResponse

func GetPokeApiLocations(next bool) {
	if len(PokeMap.Next) == 0 {
		PokeMap.Next = "https://pokeapi.co/api/v2/location"
	}
	if len(PokeMap.Previous) == 0 {
		PokeMap.Previous = "https://pokeapi.co/api/v2/location"
	}
	var url string = PokeMap.Previous
	if next {
		url = PokeMap.Next
	}
	resp, err := Request(url)
	if err != nil {
		log.Fatal("Gokeded", err)
	}
	unmarshalErr := json.Unmarshal(resp, &PokeMap)
	if unmarshalErr != nil {
		log.Fatal("Gokeded", err)
	}
	for _, result := range PokeMap.Results {
		fmt.Println(result.Name)
	}

}

func GetPokemonsInLocation(arg string) {

}
