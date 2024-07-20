package handler

import (
	"main.go/internal/domain"
	"main.go/internal/usecase"
)

// make an interface for player handler
type PlayerHandlerInterface interface {
	PlayerAdd
	PlayerUpdate
	PlayerDelete
	PlayerView
	PlayerAddPokemon
	PlayerLogin
	PlayerViewTheirPokemon
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
	PlayerAddPokemon(playerUsn string, pokemon domain.Pokemon) error
}
type PlayerLogin interface {
	PlayerLogin(playerUsn string) (*domain.Player, error)
}
type PlayerViewTheirPokemon interface {
	PlayerViewTheirPokemon(playerUsn string) error
}

// make a struct
type PlayerHandler struct {
	PlayerUsecase usecase.PlayerUsecaseInterface
}

// make a connection to the usecase
func NewPlayerHandler(playerUsecase usecase.PlayerUsecaseInterface) PlayerHandlerInterface {
	return PlayerHandler{
		PlayerUsecase: playerUsecase,
	}
}

// implement the interface for player handler
func (h PlayerHandler) PlayerAdd(player domain.Player) error {
	// validate the input first
	if err := validate.Struct(player); err != nil {
		return err
	}
	return h.PlayerUsecase.PlayerAdd(player)
	// err := h.PlayerUsecase.PlayerAdd(player)
	// if err != nil {
	// 	return err
	// }
	// return nil
}
func (h PlayerHandler) PlayerUpdate(player domain.Player) error {
	//validate first
	if err := validate.Struct(player); err != nil {
		return err
	}
	err := h.PlayerUsecase.PlayerUpdate(player)
	if err != nil {
		return err
	}
	return nil
}
func (h PlayerHandler) PlayerDelete(playerId int) error {
	err := h.PlayerUsecase.PlayerDelete(playerId)
	if err != nil {
		return err
	}
	return nil
}
func (h PlayerHandler) PlayerView() error {
	err := h.PlayerUsecase.PlayerView()
	if err != nil {
		return err
	}
	return nil
}
func (h PlayerHandler) PlayerAddPokemon(playerUsn string, pokemon domain.Pokemon) error {
	err := h.PlayerUsecase.PlayerAddPokemon(playerUsn, pokemon)
	if err != nil {
		return err
	}
	return nil
}
func (h PlayerHandler) PlayerLogin(playerUsn string) (*domain.Player, error) {
	player, err := h.PlayerUsecase.PlayerLogin(playerUsn)
	if err != nil {
		return nil, err
	}
	return player, nil
}

func (h PlayerHandler) PlayerViewTheirPokemon(playerUsn string) error {
	err := h.PlayerUsecase.PlayerViewTheirPokemon(playerUsn)
	if err != nil {
		return err
	}
	return nil
}
