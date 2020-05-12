// Package collection is used for gathering collection resources
package collection

import (
	"fmt"

	"github.com/zmb3/spotify"
)

func AlbumFmt(album spotify.SavedAlbum) string {
	return fmt.Sprintf("%s - %s", album.Artists[0].Name, album.Name)
}

type albumsClient interface {
	CurrentUsersAlbums() (*spotify.SavedAlbumPage, error)
}

func Albums(client albumsClient) (output []spotify.SavedAlbum, err error) {
	albums, err := client.CurrentUsersAlbums()
	if err != nil {
		return
	}

	return albums.Albums, nil
}
