package main

import (
	"time"

	"main.go/internal/domain"
	"main.go/internal/handler"
	"main.go/internal/repository"
	"main.go/internal/usecase"
)

func main() {
	repo := repository.NewPokemonRepo()
	uc := usecase.NewPokemonUsecase(repo)
	h := handler.NewPokemonHandler(uc)

	pokemon := domain.Pokemon{
		Name:           "Pikachu",
		Type:           "Lightning",
		CatchRate:      70,
		IsRare:         false,
		RegisteredDate: time.Now(),
	}
	pokemon2 := domain.Pokemon{
		Name:           "Pikachu2",
		Type:           "Lightning",
		CatchRate:      70,
		IsRare:         false,
		RegisteredDate: time.Now(),
	}

	h.PokemonAdd(pokemon)
	h.PokemonAdd(pokemon2)
	h.PokemonView()
	h.PokemonDelete(2)
	h.PokemonView()

}
