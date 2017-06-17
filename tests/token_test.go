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
	t.Run("should generate the right URL", func(t *testing.T) {
		v := make(url.Values)
		v.Set("command", "request")

		req, _ := client.NewRequest("api.php", v)

		if req.URL.String() != "https://opentdb.com/api.php?command=request" {
			t.Fail()
		}
	})
}
