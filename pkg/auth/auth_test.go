package auth

import (
	"errors"
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

type errReader struct{}

func (e errReader) Read([]byte) (int, error) {
	return 0, errors.New("test error")
}

func TestLoadTokenReadErr(t *testing.T) {
	_, err := LoadToken(errReader{})
	assert.NotNil(t, err)
}

func TestLoadTokenJSONErr(t *testing.T) {
	r := strings.NewReader("")
	_, err := LoadToken(r)

	assert.NotNil(t, err)
}
