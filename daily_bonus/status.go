package dailybonus

// A Status holds daily bonus status information for a single user
type Status struct {
	// IsAvailable is true if the bonus is available
	IsAvailable bool
	// SecondsUntilAvailable is the number of seconds until the next daily bonus is available
	SecondsUntilAvailable int32
	// Streak is the current daily bonus streak of the user
	Streak int32
	// WheelValues holds the ordered collection of wheel award values to be played next
	WheelValues []int64
}
