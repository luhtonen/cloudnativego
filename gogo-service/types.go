package main

import "time"

type newMatchResponse struct {
	ID          string `json:"id"`
	StartedAt   int64  `json:"started_at"`
	GridSize    int    `json:"gridsize"`
	PlayerWhite string `json:"playerWhite"`
	PlayerBlack string `json:"playerBlack"`
	Turn        int    `json:"turn,omitempty"`
}

func (m *newMatchResponse) copyMatch(match Match) {
	m.ID = match.ID
	m.StartedAt = match.StartTime.Unix()
	m.GridSize = match.GridSize
	m.PlayerWhite = match.PlayerWhite
	m.PlayerBlack = match.PlayerBlack
	m.Turn = match.TurnCount
}

type newMatchRequest struct {
	GridSize    int    `json:"gridsize"`
	PlayerWhite string `json:"playerWhite"`
	PlayerBlack string `json:"playerBlack"`
}

func (request newMatchRequest) isValid() (valid bool) {
	valid = true
	if request.GridSize != 19 && request.GridSize != 13 && request.GridSize != 9 {
		valid = false
	}
	if request.PlayerWhite == "" {
		valid = false
	}
	if request.PlayerBlack == "" {
		valid = false
	}
	return valid
}

// GameBoard is a two-dimensional array of stones.
type GameBoard struct {
	Positions [][]byte
}

// Match represents the state of an in-progress game of Go.
type Match struct {
	TurnCount   int
	GridSize    int
	ID          string
	StartTime   time.Time
	GameBoard   GameBoard
	PlayerBlack string
	PlayerWhite string
}

type matchRepository interface {
	addMatch(match Match) (err error)
	getMatches() (matches []Match, err error)
	getMatch(id string) (match Match, err error)
	updateMatch(id string, match Match) (err error)
}
