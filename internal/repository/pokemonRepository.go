package repository

import (
	"errors"
	"fmt"

	"main.go/internal/domain"
)

// make an interface for pokemon repo
type PokemonRepoInterface interface {
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
func NewPokemonRepo() PokemonRepoInterface {
	return PokemonRepo{
		Pokemon: map[int]domain.Pokemon{},
	}
}

// implement the required functions
func (repo PokemonRepo) PokemonAdd(pokemon *domain.Pokemon) error {
	if len(repo.Pokemon) == 0 {
		pokemon.ID = 1
	} else {
		pokemon.ID = repo.Pokemon[len(repo.Pokemon)].ID + 1
	}
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
	if len(repo.Pokemon) == 0 || repo.Pokemon == nil {
		err := errors.New("POKEMON LIST IS EMPTY 🤬🤬🤬")
		fmt.Println(err)
	}
	//print the pokemon list
	for _, value := range repo.Pokemon {
		fmt.Printf("%d %s %s %f %t ", value.ID, value.Name, value.Type, value.CatchRate, value.IsRare)
		fmt.Println(value.RegisteredDate.Format("2006-01-02"))
	}
	fmt.Println()
	return nil
}
