package userdata

// A User holds UserData information for a single player
type User struct {
	Coins  int64
	ID     string
	Rating *Rating
}

// Rating holds Glicko2 rating data
type Rating struct {
	Value      float64
	Deviation  float64
	Volatility float64
}

// A MatchResult holds match results for a pair of players
type MatchResult struct {
	UserID1 string
	UserID2 string
	Score   MatchScore
}
