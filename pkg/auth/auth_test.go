package auth

import (
	"bufio"
	"bytes"
	"errors"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
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

type errReadWriter struct{}

func (e errReadWriter) Read([]byte) (int, error) {
	return 0, errors.New("test error")
}

func (e errReadWriter) Write([]byte) (int, error) {
	return 0, errors.New("test error")
}

func TestLoadTokenReadErr(t *testing.T) {
	_, err := LoadToken(errReadWriter{})
	assert.NotNil(t, err)
}

func TestLoadTokenJSONErr(t *testing.T) {
	r := strings.NewReader("")
	_, err := LoadToken(r)

	assert.NotNil(t, err)
}

func TestSaveToken(t *testing.T) {
	buf := &bytes.Buffer{}
	token := oauth2.Token{}

	err := SaveToken(token, buf)
	assert.Nil(t, err)

	j, err := ioutil.ReadFile("testdata/token.json")
	assert.Nil(t, err)
	assert.Equal(t, string(j), buf.String())
}

type errJSON struct{}

func (e errJSON) MarshalJSON() ([]byte, error) {
	return nil, errors.New("Marshal error")
}

func TestSaveTokenJsonErr(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	token := errJSON{}

	err := SaveToken(&token, w)
	assert.NotNil(t, err)
}

func TestSaveTokenWriteErr(t *testing.T) {
	token := oauth2.Token{}
	w := errReadWriter{}

	err := SaveToken(&token, w)
	assert.NotNil(t, err)
}
