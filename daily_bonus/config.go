package dailybonus

import (
	"github.com/mattwhip/icenine-services/math/sampling"
	"github.com/pkg/errors"
)

// Config holds daily bonus config
type Config struct {
	ResetSeconds       int32
	StreakBreakSeconds int32
	Wheels             *Wheels
}

// GetWheel retrieves the wheel for a specific daily bonus streak values
func (c Config) GetWheel(streak int) *Wheel {
	switch streak {
	case 0:
		return c.Wheels.Streak0
	case 1:
		return c.Wheels.Streak1
	case 2:
		return c.Wheels.Streak2
	case 3:
		return c.Wheels.Streak3
	case 4:
		return c.Wheels.Streak4
	case 5:
		return c.Wheels.Streak5
	}
	return nil
}

// Validate returns nil if config is valid or an error describing why it is invalid
func (c Config) Validate() error {
	if c.ResetSeconds <= 0 {
		return errors.New("ResetSeconds must be positive")
	}
	if c.StreakBreakSeconds <= 0 {
		return errors.New("StreakBreakSeconds must be positive")
	}
	if c.Wheels == nil {
		return errors.New("wheels cannot be null")
	}
	if err := c.Wheels.Validate(); err != nil {
		return errors.Wrap(err, "failed to validate wheels")
	}
	return nil
}

// Wheels holds configuration for all wheels
type Wheels struct {
	Streak0 *Wheel `json:"streak_0"`
	Streak1 *Wheel `json:"streak_1"`
	Streak2 *Wheel `json:"streak_2"`
	Streak3 *Wheel `json:"streak_3"`
	Streak4 *Wheel `json:"streak_4"`
	Streak5 *Wheel `json:"streak_5"`
}

// Validate returns nil if the Wheels are valid or an error explaining why they are invalid
func (w *Wheels) Validate() error {
	if w.Streak0 == nil || w.Streak1 == nil {
		return errors.New("must include streak_0 - streak_5 configuration")
	}
	if err := w.Streak0.Validate(); err != nil {
		return errors.Wrap(err, "streak_0 wheel is invalid")
	}
	if err := w.Streak1.Validate(); err != nil {
		return errors.Wrap(err, "streak_1 wheel is invalid")
	}
	if err := w.Streak2.Validate(); err != nil {
		return errors.Wrap(err, "streak_2 wheel is invalid")
	}
	if err := w.Streak3.Validate(); err != nil {
		return errors.Wrap(err, "streak_3 wheel is invalid")
	}
	if err := w.Streak4.Validate(); err != nil {
		return errors.Wrap(err, "streak_4 wheel is invalid")
	}
	if err := w.Streak5.Validate(); err != nil {
		return errors.Wrap(err, "streak_5 wheel is invalid")
	}
	return nil
}

// Wheel holds configuration for a single wheel
type Wheel struct {
	Slices     []Slice              `json:"slices"`
	Population *sampling.Population `json:"-"`
}

// InitializePopulation initializes the wheel's sampling population based on its Slices
func (w *Wheel) InitializePopulation() error {
	items := make([]interface{}, 0, len(w.Slices))
	weights := make([]int64, 0, len(w.Slices))
	for _, slice := range w.Slices {
		items = append(items, slice.Value)
		weights = append(weights, slice.Weight)
	}
	var err error
	w.Population, err = sampling.NewPopulation(items, weights)
	if err != nil {
		return errors.Wrap(err, "failed to create population")
	}
	return nil
}

// GetAwardValues retrieves the ordered award values on the wheel
func (w *Wheel) GetAwardValues() []int64 {
	values := make([]int64, 0, len(w.Slices))
	for _, v := range w.Slices {
		values = append(values, v.Value)
	}
	return values
}

// Validate returns nil if the wheel is valid or an error describing why it is invalid
func (w *Wheel) Validate() error {
	if len(w.Slices) < 2 {
		return errors.New("wheel must have at least two slices")
	}
	for _, slice := range w.Slices {
		if slice.Value < 1 {
			return errors.New("slice award values must be positive")
		}
		if slice.Weight < 1 {
			return errors.New("slice weights must be positive")
		}
	}
	return nil
}

// Slice holds configuration for a single wheel slice
type Slice struct {
	Value  int64 `json:"value"`
	Weight int64 `json:"weight"`
}
