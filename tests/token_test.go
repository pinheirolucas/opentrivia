package tests

import (
	"testing"
)

func TestCreate(t *testing.T) {
	t.Parallel()

	token, err := client.Token.Create()
	if err != nil {
		t.Error(err)
	}

	if token == "" {
		t.Error("Create method returned an empty token")
		return
	}
}

func TestRefresh(t *testing.T) {
	t.Parallel()

	token, err := client.Token.Create()
	if err != nil {
		panic("something went wrong on token creation")
	}

	t.Run("expect the refreshed token to be equals to the original", func(t *testing.T) {
		t.Parallel()

		resultToken, err := client.Token.Refresh(token)
		if err != nil {
			t.Error(err)
		}

		if resultToken != token {
			t.Errorf("Expected token %s, got %s", token, resultToken)
		}
	})
}
