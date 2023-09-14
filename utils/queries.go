package utils

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/AtinAgnihotri/gokedex/types"
)

var PokeLocation types.PokeLocationsListResponse
var Pokedex map[string]types.Pokemon

func GetPokeApiLocations(next bool) {
	if len(PokeLocation.Next) == 0 {
		PokeLocation.Next = "https://pokeapi.co/api/v2/location"
	}
	if len(PokeLocation.Previous) == 0 {
		PokeLocation.Previous = "https://pokeapi.co/api/v2/location"
	}
	var url string = PokeLocation.Previous
	if next {
		url = PokeLocation.Next
	}
	resp, err := Request(url)
	if err != nil {
		log.Fatal("Gokeded", err)
	}
	unmarshalErr := json.Unmarshal(resp, &PokeLocation)
	if unmarshalErr != nil {
		log.Fatal("Gokeded", unmarshalErr)
	}
	fmt.Println("\n+==== Map ====+")
	for _, result := range PokeLocation.Results {
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
			if !CheckIfExists[string](pokemons, pokemon) {
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

func GetPokemon(pokemonName string) (types.Pokemon, error) {
	var pokemon types.Pokemon
	resp, err := Request(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%v", pokemonName))
	if err != nil {
		return pokemon, err
	}
	unmarshallError := json.Unmarshal(resp, &pokemon)
	if unmarshallError != nil {
		return pokemon, unmarshallError
	}
	return pokemon, nil
}
