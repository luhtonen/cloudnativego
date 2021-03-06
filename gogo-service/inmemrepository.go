package main

import (
	"errors"
	"strings"
)

type inMemoryMatchRepository struct {
	matches []Match
}

// NewRepository creates a new in-memory match repository
func newInMemoryRepository() *inMemoryMatchRepository {
	repo := &inMemoryMatchRepository{}
	repo.matches = []Match{}
	return repo
}

func (repo *inMemoryMatchRepository) addMatch(match Match) (err error) {
	repo.matches = append(repo.matches, match)
	return err
}

func (repo *inMemoryMatchRepository) getMatches() (matches []Match, err error) {
	matches = repo.matches
	return
}

func (repo *inMemoryMatchRepository) getMatch(id string) (match Match, err error) {
	found := false
	for _, target := range repo.matches {
		if strings.Compare(target.ID, id) == 0 {
			match = target
			found = true
		}
	}
	if !found {
		err = errors.New("Could not find match in repository")
	}
	return match, err
}

func (repo *inMemoryMatchRepository) updateMatch(id string, match Match) (err error) {
	found := false
	for k, v := range repo.matches {
		if strings.Compare(v.ID, id) == 0 {
			repo.matches[k] = match
			found = true
		}
	}
	if !found {
		err = errors.New("Could not find match in repository")
	}
	return
}
