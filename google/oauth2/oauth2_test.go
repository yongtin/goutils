package oauth2

import (
	"testing"

	photoslibrary "github.com/gphotosuploader/googlemirror/api/photoslibrary/v1"
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
)

func TestGetOCConfig(t *testing.T) {
	var oc *oauth2.Config
	var err error
	var configPath = "/Users/yongtin/Downloads/client_secret_655962567107-hk3a1a2a5jqh0cnl537tc3tg93sdd5mn.apps.googleusercontent.com.json"
	oc, err = GetoAuth2ConfigFromJSON("/tmp", photoslibrary.PhotoslibraryAppendonlyScope)
	assert.NotNil(t, err)
	oc, err = GetoAuth2ConfigFromJSON("/tmp/permissionproblem.txt", photoslibrary.PhotoslibraryAppendonlyScope)
	assert.NotNil(t, err)
	oc, err = GetoAuth2ConfigFromJSON("/tpm", photoslibrary.PhotoslibraryAppendonlyScope)
	assert.NotNil(t, err)

	// os.Remove("token.json")
	// oc, err = GetoAuth2ConfigFromJSON(configPath, photoslibrary.PhotoslibraryAppendonlyScope)
	// _, err = GetHTTPClient(oc)
	// assert.Nil(t, err)
	oc, err = GetoAuth2ConfigFromJSON(configPath, photoslibrary.PhotoslibraryReadonlyScope)
	_, err = GetHTTPClient(oc)
	assert.Nil(t, err)

}
