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
package cmd

import (
	"fmt"

	"github.com/shanedabes/spu/pkg/auth"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get endpoint resources",
	Long: `Get resources from the currently authenticated user's library.

Alternatively, spotify IDs can be given to retrieve information on specific resources.`,
}

// getAlbumsCmd represents the albums command
var getAlbumsCmd = &cobra.Command{
	Use:   "albums",
	Short: "Get user albums",
	Long: `Retrieve all of the albums in the current user's library.

Alternatively, pass IDs to retrieve information on specific albums.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := auth.CachedClient()
		if err != nil {
			return err
		}

		albums, err := client.CurrentUsersAlbums()
		if err != nil {
			return err
		}

		for _, album := range albums.Albums {
			out := fmt.Sprintf(
				"%s - %s", album.Artists[0].Name, album.Name,
			)
			fmt.Println(out)
		}

		return nil
	},
}
