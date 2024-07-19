package repository

import (
	"errors"
	"fmt"

	"main.go/internal/domain"
)

// make an interface for pokemon repo
type PokemonInterface interface {
	PokemonAdd
	PokemonDelete
	PokemonView
}
type PokemonAdd interface {
	PokemonAdd(pokemon *domain.Pokemon) error
}
type PokemonDelete interface {
	PokemonDelete(pokemonId int) error
}
type PokemonView interface {
	PokemonView() error
}

// make a database (db db-an) using map
type PokemonRepo struct {
	Pokemon map[int]domain.Pokemon
}

// make a connection to the database
func NewPokemon() PokemonInterface {
	return PokemonRepo{
		Pokemon: map[int]domain.Pokemon{},
	}
}

// implement the required functions
func (repo PokemonRepo) PokemonAdd(pokemon *domain.Pokemon) error {
	_, exist := repo.Pokemon[pokemon.ID]
	if exist {
		return errors.New("POKEMON WITH SUCH ID IS ALREADY EXIST")
	}
	//add pokemon to the map
	repo.Pokemon[pokemon.ID] = *pokemon
	return nil
}
func (repo PokemonRepo) PokemonDelete(pokemonId int) error {
	_, exist := repo.Pokemon[pokemonId]
	if !exist {
		return errors.New("POKEMON WITH SUCH ID IS NOT EXIST")
	}
	//delete the pokemon with the input ID
	delete(repo.Pokemon, pokemonId)
	return nil
}
func (repo PokemonRepo) PokemonView() error {
	if len(repo.Pokemon) == 0 {
		return errors.New("POKEMON LIST IS EMPTY")
	}
	//print the pokemon list
	fmt.Println(repo.Pokemon)
	return nil
}
