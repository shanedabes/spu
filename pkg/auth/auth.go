package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
)

// LoadToken loads a token from cache file (or any io.Reader)
func LoadToken(r io.Reader) (t oauth2.Token, err error) {
	j, err := ioutil.ReadAll(r)
	if err != nil {
		return t, err
	}

	err = json.Unmarshal(j, &t)
	if err != nil {
		return t, err
	}

	return t, nil
}

// SaveToken writes a token (or any json) to a file (or any io.Writer)
func SaveToken(i interface{}, w io.Writer) (err error) {
	j, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return
	}

	_, err = fmt.Fprint(w, string(j)+"\n")
	if err != nil {
		return
	}

	return
}

func CachedClient() (c spotify.Client, err error) {
	userCacheDir, err := os.UserCacheDir()
	if err != nil {
		return
	}

	cache, err := os.Open(path.Join(userCacheDir, "spu", "token.json"))
	if err != nil {
		return
	}
	defer cache.Close()

	token, err := LoadToken(cache)
	if err != nil {
		return
	}

	return spotify.NewAuthenticator("", "").NewClient(&token), nil
}
