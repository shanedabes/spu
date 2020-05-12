// Package collection is used for gathering collection resources
package collection

import (
	"fmt"

	"github.com/zmb3/spotify"
)

func albumFmt(album spotify.SavedAlbum) string {
	return fmt.Sprintf("%s - %s", album.Artists[0].Name, album.Name)
}

type albumsClient interface {
	CurrentUsersAlbums() (*spotify.SavedAlbumPage, error)
}

func Albums(client albumsClient) (output []string, err error) {
	albums, err := client.CurrentUsersAlbums()
	if err != nil {
		return
	}

	output = []string{}

	for _, album := range albums.Albums {
		fmt := albumFmt(album)
		output = append(output, fmt)
	}

	return
}
