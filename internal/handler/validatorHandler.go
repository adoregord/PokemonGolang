package handler

import (
	"strings"

	"github.com/go-playground/validator"
)

// deklarasi variabel validate untuk digunakan
var validate *validator.Validate

// inisiasi atau setup validatenya
func init() {
	validate = validator.New()
	validate.RegisterValidation("noblank", func(fl validator.FieldLevel) bool {
		return strings.TrimSpace(fl.Field().String()) != ""
	})
	// validate.RegisterValidation("datetime", func(fl validator.FieldLevel) bool {
	// 	dateStr := fl.Field().String()
	// 	_, err := time.Parse("2006-01-02", dateStr)
	// 	return err == nil
	// })
}
