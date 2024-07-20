package domain

type Pokemon struct {
	ID             int     // auto generated
	Name           string  `validate:"required"`
	Type           string  `validate:"required"`
	CatchRate      float32 `validate:"gt=0,lte=100"`
	IsRare         bool    
	RegisteredDate string  `validate:"datetime"`
}
