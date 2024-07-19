package usecase

import (
	"errors"

	"main.go/internal/domain"
	"main.go/internal/repository"
)

// make an interface for pokemon usecase
type PokemonUsecaseInterface interface {
	PokemonAdd
	PokemonDelete
	PokemonView
}
type PokemonAdd interface {
	PokemonAdd(pokemon domain.Pokemon) error
}
type PokemonDelete interface {
	PokemonDelete(pokemonId int) error
}
type PokemonView interface {
	PokemonView() error
}

// buat struct untuk koneksi ke repository
type PokemonUsecase struct {
	PokemonRepo repository.PokemonRepoInterface
}

// make a connection to the pokemon repository
func NewPokemonUsecase(pokemonRepo repository.PokemonRepoInterface) PokemonUsecaseInterface {
	return PokemonUsecase{
		PokemonRepo: pokemonRepo,
	}
}

// implement the interface
func (uc PokemonUsecase) PokemonAdd(pokemon domain.Pokemon) error {
	err := uc.PokemonRepo.PokemonAdd(&pokemon)
	if err != nil {
		return errors.New("KAYAKNYA ADA ERROR: " + err.Error())
	}
	return nil
}
func (uc PokemonUsecase) PokemonDelete(pokemonId int) error {
	err := uc.PokemonRepo.PokemonDelete(pokemonId)
	if err != nil {
		return errors.New("KAYAKNYA ADA ERROR: " + err.Error())
	}
	return nil
}
func (uc PokemonUsecase) PokemonView() error {
	err := uc.PokemonRepo.PokemonView()
	if err != nil {
		return errors.New("KAYAKNYA ADA ERROR: " + err.Error())
	}
	return nil
}
