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

	pokemon := []domain.Pokemon{
		{Name: "Pikachu", Type: "Lightning", CatchRate: 70, IsRare: false, RegisteredDate: time.Now()},
		{Name: "Bulbasaur", Type: "Plant", CatchRate: 50, IsRare: false, RegisteredDate: time.Now()},
		{Name: "Charmander", Type: "Fire", CatchRate: 50, IsRare: false, RegisteredDate: time.Now()},
		{Name: "Rattata", Type: "Normal", CatchRate: 80, IsRare: false, RegisteredDate: time.Now()},
		{Name: "Ditto", Type: "Normal", CatchRate: 30, IsRare: true, RegisteredDate: time.Now()},
		{Name: "Mew Two", Type: "Psychic", CatchRate: 0.001, IsRare: true, RegisteredDate: time.Now()},
		{Name: "Dialga", Type: "Steel/Dragon", CatchRate: 0.001, IsRare: true, RegisteredDate: time.Now()},
		{Name: "Arceus", Type: "Normal", CatchRate: 0.0000000000001, IsRare: true, RegisteredDate: time.Now()},
	}

	pokemon2 := domain.Pokemon{
		ID: 2, 
		Name: "Arceussss", 
		Type: "Normal", 
		CatchRate: 100, 
		IsRare: true, 
		RegisteredDate: time.Now(),
	}

	for _, value := range pokemon {
		h.PokemonAdd(value)
	}
	h.PokemonView()
	h.PokemonDelete(1)
	h.PokemonUpdate(pokemon2)
	h.PokemonView()

}
