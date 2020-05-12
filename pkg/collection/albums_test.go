package collection

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zmb3/spotify"
)

var testAlbum = spotify.SavedAlbum{
	FullAlbum: spotify.FullAlbum{
		SimpleAlbum: spotify.SimpleAlbum{
			Artists: []spotify.SimpleArtist{
				{
					Name: "test",
				},
			},
			Name: "test",
		},
	},
}

func TestAlbumFmt(t *testing.T) {
	expected := "test - test"
	got := albumFmt(testAlbum)

	assert.Equal(t, expected, got)
}

type mockAlbumsClient struct{}

func (m *mockAlbumsClient) CurrentUsersAlbums() (*spotify.SavedAlbumPage, error) {
	return &spotify.SavedAlbumPage{
		Albums: []spotify.SavedAlbum{
			testAlbum,
		},
	}, nil
}

func TestAlbums(t *testing.T) {
	expected := []string{
		"test - test",
	}
	got, err := Albums(&mockAlbumsClient{})

	assert.Nil(t, err)
	assert.Equal(t, expected, got)
}

type mockAlbumsErrorClient struct{}

func (m *mockAlbumsErrorClient) CurrentUsersAlbums() (*spotify.SavedAlbumPage, error) {
	return &spotify.SavedAlbumPage{}, errors.New("test")
}

func TestAlbumsGetAlbumsError(t *testing.T) {
	_, err := Albums(&mockAlbumsErrorClient{})
	assert.NotNil(t, err)
}
