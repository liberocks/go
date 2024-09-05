package helpers

import (
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

var vOnce sync.Once
var v *validator.Validate

// GetValidator is responsible for returning a single instance of the validator.
func GetValidator() *validator.Validate {
	vOnce.Do(func() {
		log.Info().Msg("Validator initialized.")
		v = validator.New()
	})

	return v
}