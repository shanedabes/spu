/*
Copyright © 2020 Shane Donohoe <shane@isda.best>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package cmd contains code for the spu cli
package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/shanedabes/spu/pkg/auth"
	"github.com/spf13/cobra"
	"github.com/zmb3/spotify"
)

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate with the spotify API",
	Long: `Authenticate with the spotify API using client and secret variables. These can be provided using environment variables, flags or from the config file.

The generated token will be saved to cache to prevent the need to run this command again.`,
	RunE: authFunc,
}

const (
	state = "abc123"
)

func authFunc(cmd *cobra.Command, args []string) error {
	userCacheDir, err := os.UserCacheDir()
	if err != nil {
		return err
	}

	cacheDir := path.Join(userCacheDir, "spu")

	err = os.MkdirAll(cacheDir, 0700)
	if err != nil {
		return err
	}

	cache, err := os.Create(path.Join(cacheDir, "token.json"))
	if err != nil {
		return err
	}
	defer cache.Close()

	ch := make(chan string)
	http.HandleFunc("/callback", auth.NewCallbackHandler(state, ch))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request for:", r.URL.String())
	})

	go http.ListenAndServe(":8080", nil)

	sp := spotify.NewAuthenticator(auth.RedirectURL, auth.Scopes...)
	url := sp.AuthURL(state)
	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)

	code := <-ch

	token, err := sp.Exchange(code)
	if err != nil {
		return err
	}

	err = auth.SaveToken(token, cache)
	if err != nil {
		return err
	}

	return nil
}
