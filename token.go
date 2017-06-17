package opentrivia

import (
	"errors"
)

// Token is the type for tokens.
type Token string

var (
	// ErrTokenEmpty is returned when the Open Trivia API has
	// returned all possible questions for the specified query.
	//
	// Resetting the Token is necessary to keep on running.
	ErrTokenEmpty = errors.New("opentrivia: token has returned all possible questions for the specified query")

	// ErrTokenNotFound is returned when the Open Trivia API
	// do not found the provided token.
	ErrTokenNotFound = errors.New("opentrivia: token does not exist")
)

// TokenService handles communication with the token related
// methods of the Open Trivia API
//
// Ref.: https://opentdb.com/api_config.php
type TokenService service

// Create returns a brand new token from Open Trivia API.
// Each token provides the guarantee that every new requested
// question was not already retrieved.
//
// By sending a token to an API Call, the API will never return
// the same question twice.
//
// If all questions for a given category has already been returned,
// the request will return an opentrivia.ErrTokenEmpty.
func (t *TokenService) Create() (Token, error) {
	return "", nil
}

// Reset refresh the provided token.
//
// If the provided token is invalid, the request will return an
// opentrivia.ErrTokenNotFound.
func (t *TokenService) Reset(token Token) (Token, error) {
	return "", nil
}

// StillValid should be used to check if the token still valid.
//
// If the provided token is not found, the request will return an
// opentrivia.ErrTokenNotFound.
func (t *TokenService) StillValid(token Token) (bool, error) {
	return false, nil
}
