package usecase

import (
	"main.go/internal/domain"
	"main.go/internal/repository"
)

// make an interface for pokemon usecase
type PlayerUsecaseInterface interface {
	PlayerAdd
	PlayerUpdate
	PlayerDelete
	PlayerView
	PlayerAddPokemon
}
type PlayerAdd interface {
	PlayerAdd(player domain.Player) error
}
type PlayerUpdate interface {
	PlayerUpdate(player domain.Player) error
}
type PlayerDelete interface {
	PlayerDelete(playerId int) error
}
type PlayerView interface {
	PlayerView() error
}
type PlayerAddPokemon interface {
	PlayerAddPokemon(playerId int, pokemon domain.Pokemon) error
}

// make a struct
type PlayerUsecase struct {
	PlayerRepo repository.PlayerRepoInterface
}

// make a connection to the repository
func NewPlayerUsecase(playerRepo repository.PlayerRepoInterface) PlayerUsecaseInterface {
	return PlayerUsecase{
		PlayerRepo: playerRepo,
	}
}

// make an implementation for player usecase
func (uc PlayerUsecase) PlayerAdd(player domain.Player) error {
	err := uc.PlayerRepo.PlayerAdd(&player)
	if err != nil{
		return err
	}
	return nil
}
func (uc PlayerUsecase) PlayerUpdate(player domain.Player) error {
	err := uc.PlayerRepo.PlayerUpdate(&player)
	if err != nil {
		return err
	}
	return nil
}
func (uc PlayerUsecase) PlayerDelete(playerId int) error {
	err := uc.PlayerRepo.PlayerDelete(playerId)
	if err != nil {
		return err
	}
	return nil
}
func (uc PlayerUsecase) PlayerView() error {
	err := uc.PlayerRepo.PlayerView()
	if err != nil {
		return err
	}
	return nil
}
func (uc PlayerUsecase) PlayerAddPokemon(playerId int, pokemon domain.Pokemon) error {
	err := uc.PlayerRepo.PlayerAddPokemon(playerId, &pokemon)
	if err != nil {
		return err
	}
	return nil
}
