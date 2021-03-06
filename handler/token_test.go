package handler

import (
	"testing"
)

func Test_newToken(t *testing.T) {
	var token = newToken(10)
	t.Logf("new token: %s", token)
}