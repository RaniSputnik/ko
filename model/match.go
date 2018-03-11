package model

type MatchStatus string

const (
	MatchStatusInvalid    = ""
	MatchStatusWaiting    = "WAITING_FOR_OPPONENT"
	MatchStatusReady      = "READY"
	MatchStatusInProgress = "IN_PROGRESS"
	MatchStatusCompleted  = "COMPLETED"
)

type Match struct {
	ID        string
	Owner     string
	Opponent  string
	BoardSize int
}

func (m Match) Status() MatchStatus {
	if m.Opponent != "" {
		return MatchStatusReady
	}
	return MatchStatusWaiting
}
