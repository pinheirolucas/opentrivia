package opentrivia

import "github.com/pkg/errors"
import "github.com/google/go-querystring/query"

type (
	// Token is the type for tokens.
	Token string

	tokenCommand string
)

const (
	tokenCommandCreate  tokenCommand = "request"
	tokenCommandRefresh tokenCommand = "reset"
)

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

type tokenOptions struct {
	Command tokenCommand `url:"command,omitempty"`
	Token   Token        `url:"token,omitempty"`
}

type tokenResponse struct {
	ResponseCode    responseCode `json:"response_code,omitempty"`
	ResponseMessage string       `json:"response_message,omitempty"`
	Token           Token        `json:"token,omitempty"`
}

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
	options := &tokenOptions{
		Command: tokenCommandCreate,
	}

	v, err := query.Values(options)
	if err != nil {
		return "", err
	}

	req, err := t.client.NewRequest(defaultTokenRoute, v)
	if err != nil {
		return "", err
	}

	var resp tokenResponse
	if _, err := t.client.Do(req, &resp); err != nil {
		return "", err
	}

	switch resp.ResponseCode {
	case responseCodeInvalidParameter:
		return "", ErrInvalidParameter
	case responseCodeTokenNotFound:
		return "", ErrTokenNotFound
	}

	return resp.Token, nil
}

// Refresh the provided token.
//
// If the provided token is invalid, the request will return an
// opentrivia.ErrTokenNotFound.
func (t *TokenService) Refresh(token Token) (Token, error) {
	options := &tokenOptions{
		Command: tokenCommandRefresh,
		Token:   token,
	}

	v, err := query.Values(options)
	if err != nil {
		return "", err
	}

	req, err := t.client.NewRequest(defaultTokenRoute, v)
	if err != nil {
		return "", err
	}

	var resp tokenResponse
	if _, err := t.client.Do(req, &resp); err != nil {
		return "", err
	}

	switch resp.ResponseCode {
	case responseCodeInvalidParameter:
		return "", ErrInvalidParameter
	case responseCodeTokenNotFound:
		return "", ErrTokenNotFound
	}

	return resp.Token, nil
}
