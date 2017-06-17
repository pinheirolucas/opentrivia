package opentrivia

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

type (
	responseCode uint8
)

const (
	responseCodeSuccess          responseCode = 0
	responseCodeNoResults        responseCode = 1
	responseCodeInvalidParameter responseCode = 2
	responseCodeTokenNotFound    responseCode = 3
	responseCodeTokenEmpty       responseCode = 4
)

const (
	defaultBaseURL    = "https://opentdb.com/"
	defaultAPIRoute   = "api.php"
	defaultTokenRoute = "api_token.php"
)

// DefaultClient is the default client for Open Trivia API.
// It is the same as calling opentrivia.NewClient(nil).
var DefaultClient = NewClient(nil)

var (
	// ErrInvalidParameter is returned when the Open Trivia API
	// identifies an invalid parameter.
	ErrInvalidParameter = errors.New("opentrivia: the provided options are not valid")

	// ErrNoResults is returned when the Open Trivia API has no
	// results.
	ErrNoResults = errors.New("opentrivia: no results were found for the provided options")
)

type service struct {
	client *Client
}

// A Client manages communication with the Open Trivia API.
type Client struct {
	client *http.Client
	common service

	// Base URL for API requests. Defaults to the public Open Trivia API.
	// BaseURL should always be especified with a trailing slash.
	BaseURL *url.URL

	// Services used for talking to different parts of the Open Trivia API.
	// TODO: Add the services.
	Question *QuestionService
	Token    *TokenService
}

// NewClient returns a new Open Trivia API client. If a nil httpClient is
// provided, http.DefaultClient will be used.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		cloned := *http.DefaultClient
		httpClient = &cloned
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{
		client:  httpClient,
		BaseURL: baseURL,
	}

	c.common.client = c
	c.Question = (*QuestionService)(&c.common)
	c.Token = (*TokenService)(&c.common)

	return c
}

// NewRequest creates an API request.
func (c *Client) NewRequest(r string, q url.Values) (*http.Request, error) {
	rel, err := url.Parse(r)
	if err != nil {
		return nil, err
	}

	rel.RawQuery = q.Encode()

	u := c.BaseURL.ResolveReference(rel)

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// Do sends an API request and returns an API response.The API response is
// decoded and stored in the value pointed to by v, or returned as an error
// if an API error has occurred.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	// Make sure to close the connection after replying to this request
	req.Close = true

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}

	if err != nil {
		return nil, errors.Wrapf(
			err,
			"opentrivia: error reading response from %s %s",
			req.Method,
			req.URL.RequestURI(),
		)
	}

	return resp, nil
}
