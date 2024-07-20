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
	PlayerLogin
	PlayerViewTheirPokemon
}
type PlayerAdd interface {
	PlayerAdd(player domain.Player) (*domain.Player, error)
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
	PlayerAddPokemon(playerUsn string, pokemon domain.Pokemon) error
}
type PlayerLogin interface {
	PlayerLogin(playerUsn string) (*domain.Player, error)
}
type PlayerViewTheirPokemon interface {
	PlayerViewTheirPokemon(playerUsn string) error
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
func (uc PlayerUsecase) PlayerAdd(player domain.Player) (*domain.Player, error) {
	value ,err := uc.PlayerRepo.PlayerAdd(&player)
	if err != nil {
		return nil, err
	}
	return value, nil
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
func (uc PlayerUsecase) PlayerAddPokemon(playerUsn string, pokemon domain.Pokemon) error {
	err := uc.PlayerRepo.PlayerAddPokemon(playerUsn, &pokemon)
	if err != nil {
		return err
	}
	return nil
}
func (uc PlayerUsecase) PlayerLogin(playerUsn string) (*domain.Player, error) {
	player, err := uc.PlayerRepo.PlayerLogin(playerUsn)
	if err != nil {
		return nil, err
	}
	return player, nil
}

func (uc PlayerUsecase) PlayerViewTheirPokemon(playerUsn string) error {
	err := uc.PlayerRepo.PlayerViewTheirPokemon(playerUsn)
	if err != nil {
		return err
	}
	return nil
}
