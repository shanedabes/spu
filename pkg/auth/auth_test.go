package auth

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadToken(t *testing.T) {
	r := strings.NewReader(`{
  "access_token": "at",
  "token_type": "Bearer",
  "refresh_token": "rt",
  "expiry": "2020-05-08T17:32:47.078715277+01:00"
}`)

	_, err := LoadToken(r)
	assert.Nil(t, err)
}
