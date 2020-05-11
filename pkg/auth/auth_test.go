package auth

import (
	"bufio"
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"path"
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

func TestDefaultCacheFileName(t *testing.T) {
	cfg := cachedClientConfig{}

	userCacheDir, err := os.UserCacheDir()
	assert.Nil(t, err)

	expected := path.Join(userCacheDir, "spu", "token.json")

	got, err := cfg.FileName()
	assert.Nil(t, err)

	assert.Equal(t, expected, got)
}

func TestSetCacheFileName(t *testing.T) {
	cfg := cachedClientConfig{}

	f := SetCacheFileName("test/token.json")
	f(&cfg)

	fn, err := cfg.FileName()
	assert.Nil(t, err)

	assert.Equal(t, fn, "test/token.json")
}

func TestCachedClient(t *testing.T) {
	testCache := path.Join("testdata", "token.json")

	_, err := CachedClient(SetCacheFileName(testCache))
	assert.Nil(t, err)
}

func TestCachedClientCfgFileNameFail(t *testing.T) {
	home := os.Getenv("HOME")
	os.Setenv("HOME", "")
	defer os.Setenv("HOME", home)

	cfg := cachedClientConfig{}
	_, err := cfg.FileName()
	assert.NotNil(t, err)

	_, err = CachedClient()
	assert.NotNil(t, err)
}

func TestCachedClientOpenFail(t *testing.T) {
	testCache := path.Join("testdata", "na_token.json")

	_, err := CachedClient(SetCacheFileName(testCache))
	assert.NotNil(t, err)
}

func TestCachedClientLoadTokenFail(t *testing.T) {
	_, err := CachedClient(SetCacheFileName("auth_test.go"))
	assert.NotNil(t, err)
}
