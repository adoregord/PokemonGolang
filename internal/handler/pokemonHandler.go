package handler

import (
	"errors"

	"main.go/internal/domain"
	"main.go/internal/usecase"
)

// make an interface for pokemon handler
type PokemonHandlerInterface interface {
	PokemonAdd
	PokemonDelete
	PokemonView
}
type PokemonAdd interface {
	PokemonAdd(pokemon domain.Pokemon) error
}
type PokemonDelete interface {
	PokemonDelete(pokemoId int) error
}
type PokemonView interface {
	PokemonView() error
}

// buat struct
type PokemonHandler struct {
	PokemonUsecase usecase.PokemonUsecaseInterface
}

// make a connection to pokemon usecase
func NewPokemonHandler(pokemonUsecase usecase.PokemonUsecaseInterface) PokemonHandlerInterface {
	return PokemonHandler{
		PokemonUsecase: pokemonUsecase,
	}
}

// implement the interface
func (h PokemonHandler) PokemonAdd(pokemon domain.Pokemon) error {
	err := h.PokemonUsecase.PokemonAdd(pokemon)
	if err != nil {
		return errors.New("KAYAKNYA ADA ERROR DEH: " + err.Error())
	}
	return nil
}
func (h PokemonHandler) PokemonDelete(pokemonId int) error {
	err := h.PokemonUsecase.PokemonDelete(pokemonId)
	if err != nil {
		return errors.New("KAYAKNYA ADA ERROR DEH: " + err.Error())
	}
	return nil
}
func (h PokemonHandler) PokemonView() error {
	err := h.PokemonUsecase.PokemonView()
	if err != nil {
		return errors.New("KAYAKNYA ADA ERROR DEH: " + err.Error())
	}
	return nil
}
