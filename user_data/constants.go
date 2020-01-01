package userdata

// MatchScore holds win/loss/draw data
type MatchScore float64

const (
	// MatchScoreWin is the win value
	MatchScoreWin MatchScore = 1.0
	// MatchScoreDraw is the draw value
	MatchScoreDraw MatchScore = 0.5
	// MatchScoreLoss is the loss value
	MatchScoreLoss MatchScore = 0.0
)
