package main

import (
	"encoding/json"
	"github.com/pborman/uuid"
	"github.com/unrolled/render"
	"io/ioutil"
	"net/http"
	"time"
)

// NewMatch creates a new match
func NewMatch(size int, playerBlackName, playerWhiteName string) Match {
	result := Match{}
	result.ID = uuid.New()
	result.StartTime = time.Now()
	result.GameBoard = newBoard(size)
	result.TurnCount = 0
	result.GridSize = size
	result.PlayerBlack = playerBlackName
	result.PlayerWhite = playerWhiteName

	return result
}

// NewBoard creates a new gameboard of a given size. Gameboards must always be square.
func newBoard(size int) GameBoard {
	outBoard := GameBoard{}
	a := make([][]byte, size)
	for i := range a {
		a[i] = make([]byte, size)
	}
	outBoard.Positions = a
	return outBoard
}

func createMatchHandler(formatter *render.Render, repo matchRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		payload, _ := ioutil.ReadAll(req.Body)
		var newMatchRequest newMatchRequest
		err := json.Unmarshal(payload, &newMatchRequest)
		if err != nil {
			formatter.Text(w, http.StatusBadRequest, "Failed to parse match request")
			return
		}
		if !newMatchRequest.isValid() {
			formatter.Text(w, http.StatusBadRequest, "Invalid new match request")
			return
		}

		newMatch := NewMatch(newMatchRequest.GridSize, newMatchRequest.PlayerBlack, newMatchRequest.PlayerWhite)
		repo.addMatch(newMatch)
		var mr newMatchResponse
		mr.copyMatch(newMatch)
		w.Header().Add("Location", "/matches/"+newMatch.ID)
		formatter.JSON(w, http.StatusCreated, &mr)
	}
}
