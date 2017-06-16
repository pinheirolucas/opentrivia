package opentrivia

import (
	"fmt"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL    = "https://opentdb.com/"
	defaultAPIRoute   = "api.php"
	defaultTokenRoute = "api_token.php"
)

var (
	apiURI, _   = url.Parse(fmt.Sprintf("%s%s", defaultBaseURL, defaultAPIRoute))
	tokenURI, _ = url.Parse(fmt.Sprintf("%s%s", defaultBaseURL, defaultTokenRoute))

	// DefaultClient is the default client for Open Trivia API.
	// It is the same as calling opentrivia.NewClient(nil).
	DefaultClient = NewClient(nil)
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

	// Base Token URL for token requests. Defaults to the public Open Trivia Token.
	// BaseTokenURL should always be especified with a trailing slash.
	BaseTokenURL *url.URL

	// Services used for talking to different parts of the Open Trivia API.
	// TODO: Add the services.
	Question *QuestionService
	Token    *TokenService
}

// NewClient returns a new Open Trivia API client. If a nil httpClient is
// provided, http.DefaultClient will be used.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL := apiURI
	baseTokenURL := tokenURI

	c := &Client{
		client:       httpClient,
		BaseURL:      baseURL,
		BaseTokenURL: baseTokenURL,
	}

	c.common.client = c
	c.Question = (*QuestionService)(&c.common)
	c.Token = (*TokenService)(&c.common)

	return c
}

// NewRequest creates an API request.
func (c *Client) NewRequest() (*http.Request, error) {
	return nil, nil
}
