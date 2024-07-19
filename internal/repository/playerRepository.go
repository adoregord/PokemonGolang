package repository

import (
	"errors"
	"fmt"

	"main.go/internal/domain"
)

// make interface for player repository
type PlayerRepoInterface interface {
	PlayerAdd
	PlayerUpdate
	PlayerDelete
	PlayerView
	PlayerAddPokemon
}
type PlayerAdd interface {
	PlayerAdd(player *domain.Player) error
}
type PlayerUpdate interface {
	PlayerUpdate(player *domain.Player) error
}
type PlayerDelete interface {
	PlayerDelete(playerId int) error
}
type PlayerView interface {
	PlayerView() error
}
type PlayerAddPokemon interface {
	PlayerAddPokemon(playerId int, pokemon *domain.Pokemon) error
}

// make a database (db db-an) player using map
type PlayerRepo struct {
	Player map[int]domain.Player
}


// make a connection to the database
func NewPlayerRepo() PlayerRepoInterface {
	return PlayerRepo{
		Player: map[int]domain.Player{},
	}
}

// make an implementation for player repo
func (repo PlayerRepo) PlayerAdd(player *domain.Player) error {
	if len(repo.Player) == 0 {
		player.ID = 1
	} else {
		player.ID = repo.Player[len(repo.Player)].ID + 1
	}
	//add player to the map (database)
	repo.Player[player.ID] = *player
	return nil
}

// player update for updating the player
func (repo PlayerRepo) PlayerUpdate(player *domain.Player) error {
	_, exist := repo.Player[player.ID]
	if !exist {
		return errors.New("PLAYER WITH SUCH ID IS NOT EXIST 🤬🤬")
	}
	repo.Player[player.ID] = *player
	return nil
}

// PlayerAddPokemon implements PlayerRepoInterface for updating the player's pokemon list
func (repo PlayerRepo) PlayerAddPokemon(playerId int, pokemon *domain.Pokemon) error {
	for index, value := range repo.Player{
		if value.ID == playerId{
			value.ListPokemon = append(value.ListPokemon, *pokemon)
			repo.Player[index] = value
			return nil
		}
	}
	return errors.New("PLAYER WITH SUCH ID IS NOT EXIST 🤬🤬🤬")
}

func (repo PlayerRepo) PlayerDelete(playerId int) error {
	_, exist := repo.Player[playerId]
	if !exist{
		return errors.New("PLAY WITH SUCH ID IS NOT EXIST 🤬🤬🤬")
	} 
	//delete the player with the input ID
	delete(repo.Player, playerId)
	return nil
}

func (repo PlayerRepo) PlayerView() error {
	if len(repo.Player) == 0 || repo.Player == nil{
		return errors.New("POKEMON LIST IS EMPTY 🤬🤬🤬")
	}

	for _, value := range repo.Player{
		fmt.Printf("Player %s (%d): \nPokemons: \n", value.UserName, value.ID)
		for index, value := range value.ListPokemon{
			fmt.Printf("%d %s - %s - %t - %f\n", index+1, value.Name, value.Type, value.IsRare, value.CatchRate)
		}
	}
	fmt.Println()
	return nil
}
