package utils

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/AtinAgnihotri/gokedex/types"
)

var PokeMap types.PokeLocationsListResponse

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
		log.Fatal("Gokeded", unmarshalErr)
	}
	fmt.Println("\n+==== Map ====+")
	for _, result := range PokeMap.Results {
		fmt.Println(result.Name)
	}
	fmt.Println()
}

func getPokemonLocationAreas(location string) (types.PokeLocationAreaListResponse, error) {
	var PokeLocationList types.PokeLocationAreaListResponse
	resp, err := Request(fmt.Sprintf("https://pokeapi.co/api/v2/location/%v", location))
	if err != nil {

		return PokeLocationList, err
	}
	unmarshalErr := json.Unmarshal(resp, &PokeLocationList)
	if unmarshalErr != nil {

		return PokeLocationList, unmarshalErr
	}
	return PokeLocationList, nil
}

func getPokemonsInArea(url string) (types.PokeLocationAreaResponse, error) {
	var PokemonInAreas types.PokeLocationAreaResponse
	resp, err := Request(url)
	if err != nil {
		return PokemonInAreas, err
	}
	unmarshalError := json.Unmarshal(resp, &PokemonInAreas)
	if unmarshalError != nil {
		return PokemonInAreas, err
	}
	return PokemonInAreas, nil
}

func checkIfExists[T comparable](arr []T, val T) bool {
	exists := false
	for _, item := range arr {
		if item == val {
			exists = true
		}
	}
	return exists
}

func GetPokemonsInLocation(location string) {
	defer fmt.Println()
	var pokemons []string
	areaData, err := getPokemonLocationAreas(location)
	if err != nil {
		fmt.Println(fmt.Sprintf("Couldn't find pokemons in: %v", location))
	}
	for _, areaDatum := range areaData.Areas {
		locationEncounters, err := getPokemonsInArea(areaDatum.URL)
		if err != nil {
			continue
		}
		for _, pokemonDatum := range locationEncounters.PokemonEncounters {
			pokemon := pokemonDatum.Pokemon.Name
			if !checkIfExists[string](pokemons, pokemon) {
				pokemons = append(pokemons, pokemon)
			}
		}
	}
	if len(pokemons) == 0 {
		fmt.Println(fmt.Sprintf("\nNo pokemons encountered in %v", location))
		return
	}
	fmt.Println(fmt.Sprintf("\n+==== Pokemon Encounters in %v ====+", location))
	for _, pokemon := range pokemons {
		fmt.Println(fmt.Sprintf(" - %v", pokemon))
	}

}
