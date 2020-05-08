package auth

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"golang.org/x/oauth2"
)

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
