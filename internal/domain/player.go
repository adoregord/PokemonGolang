package domain

type Player struct {
	ID          int       //auto generated
	UserName    string    `validate:"required,noblank"`
	ListPokemon []Pokemon `validate:"dive"`
	IsAdmin bool
}
