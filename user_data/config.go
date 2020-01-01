package userdata

import (
	"github.com/pkg/errors"
)

// Config holds user data config
type Config struct {
	InitialCoins            int64
	InitialRating           float64
	InitialRatingDeviation  float64
	InitialRatingVolatility float64
}

// Validate returns nil if config is valid or an error describing why it is invalid
func (c Config) Validate() error {
	if c.InitialCoins < 0 {
		return errors.New("InitialCoins must be greater than or equal to zero")
	}
	return nil
}
