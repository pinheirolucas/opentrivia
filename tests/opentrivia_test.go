package tests

import (
	"net/url"
	"testing"

	"github.com/pinheirolucas/opentrivia"
)

var client *opentrivia.Client

func init() {
	client = opentrivia.DefaultClient
}

func TestClientNewRequest(t *testing.T) {
	t.Run("should not return an error", func(t *testing.T) {
		t.Parallel()

		_, err := client.NewRequest("api.php", make(url.Values))

		if err != nil {
			t.Errorf("No errors expected, got: %s", err)
		}
	})

	t.Run("should generate a *http.Request with the right URL", func(t *testing.T) {
		t.Parallel()

		const expectedURL = "https://opentdb.com/api.php?command=request"

		v := make(url.Values)
		v.Set("command", "request")

		req, _ := client.NewRequest("api.php", v)
		URL := req.URL.String()

		if URL != expectedURL {
			t.Errorf("Expected %s, got %s", expectedURL, URL)
		}
	})

	t.Run("should generate a *http.Request with GET method", func(t *testing.T) {
		t.Parallel()

		const expectedMethod = "GET"

		v := make(url.Values)

		req, _ := client.NewRequest("api.php", v)
		method := req.Method

		if method != expectedMethod {
			t.Errorf("Expected %s, got %s", expectedMethod, method)
		}
	})
}
