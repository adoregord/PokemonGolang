package main

import (
	"fmt"
	"time"

	"main.go/internal/domain"
	"main.go/internal/handler"
	"main.go/internal/repository"
	"main.go/internal/usecase"
)

func main() {
	pokemonRepo := repository.NewPokemonRepo()
	pokemonUc := usecase.NewPokemonUsecase(pokemonRepo)
	pokemonH := handler.NewPokemonHandler(pokemonUc)

	playerRepo := repository.NewPlayerRepo()
	playerUc := usecase.NewPlayerUsecase(playerRepo)
	playerH := handler.NewPlayerHandler(playerUc)

	pokemon := []domain.Pokemon{
		{Name: "Pikachu", Type: "Lightning", CatchRate: 70, IsRare: false, RegisteredDate: time.Now().Format(time.RFC822)},
		{Name: "Bulbasaur", Type: "Plant", CatchRate: 50, IsRare: false, RegisteredDate: time.Now().Format(time.RFC822)},
		{Name: "Charmander", Type: "Fire", CatchRate: 50, IsRare: false, RegisteredDate: time.Now().Format(time.RFC822)},
		{Name: "Rattata", Type: "Normal", CatchRate: 80, IsRare: false, RegisteredDate: time.Now().Format(time.RFC822)},
		{Name: "Ditto", Type: "Normal", CatchRate: 30, IsRare: true, RegisteredDate: time.Now().Format(time.RFC822)},
		{Name: "Mew Two", Type: "Psychic", CatchRate: 0.001, IsRare: true, RegisteredDate: time.Now().Format(time.RFC822)},
		{Name: "Dialga", Type: "Steel/Dragon", CatchRate: 0.001, IsRare: true, RegisteredDate: time.Now().Format(time.RFC822)},
		{Name: "Arceus", Type: "Normal", CatchRate: 0.000001, IsRare: true, RegisteredDate: time.Now().Format(time.RFC822)},
	}

	pokemon2 := domain.Pokemon{
		ID:             2,
		Name:           "Arceussss",
		Type:           "Normal",
		CatchRate:      100,
		IsRare:         true,
		RegisteredDate: time.Now().Format(time.RFC822),
	}

	player := domain.Player{
		ID:       1,
		UserName: "Didi",
		ListPokemon: []domain.Pokemon{
				pokemon[1],
				pokemon[2],
		},
	}

	for _, value := range pokemon {
		pokemonH.PokemonAdd(value)
	}
	pokemonH.PokemonView()
	pokemonH.PokemonDelete(1)
	pokemonH.PokemonUpdate(pokemon2)
	pokemonH.PokemonView()

	playerH.PlayerAdd(player)
	playerH.PlayerView()
	playerH.PlayerAddPokemon(1, pokemon[7])
	playerH.PlayerView()
	playerH.PlayerDelete(1)
	fmt.Print(playerH.PlayerView())
}
