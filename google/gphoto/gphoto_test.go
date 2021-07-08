package gphoto

import (
	"fmt"
	"testing"

	photoslibrary "github.com/gphotosuploader/googlemirror/api/photoslibrary/v1"
	"github.com/stretchr/testify/assert"
	"github.com/yongtin/goutils/google/oauth2"
)

func TestUploadItem(t *testing.T) {
	var configPath = "/Users/yongtin/Downloads/client_secret_655962567107-hk3a1a2a5jqh0cnl537tc3tg93sdd5mn.apps.googleusercontent.com.json"
	oc, err := oauth2.GetoAuth2ConfigFromJSON(configPath, photoslibrary.PhotoslibraryScope)
	assert.Nil(t, err)

	httpclient, err := oauth2.GetHTTPClient(oc)
	assert.Nil(t, err)

	photoslibraryService, err := photoslibrary.New(httpclient)
	assert.Nil(t, err)

	listAlbumResponse, err := photoslibraryService.Albums.List().Do()
	assert.Nil(t, err)

	fmt.Println(listAlbumResponse.Albums)
	fmt.Println(listAlbumResponse.NextPageToken)
	assert.NotNil(t, nil)
}
