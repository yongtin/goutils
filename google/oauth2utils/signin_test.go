package oauth2utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
)

func TestGetConfig(t *testing.T) {
	var oc *oauth2.Config
	var err error
	var configPath = "/Users/yongtin/Downloads/client_secret_655962567107-hk3a1a2a5jqh0cnl537tc3tg93sdd5mn.apps.googleusercontent.com.json"
	oc, err = OAuth2ConfigFromJSON("/tmp", "")
	assert.NotNil(t, err)
	oc, err = OAuth2ConfigFromJSON("/tpm", "")
	assert.NotNil(t, err)

	oc, err = OAuth2ConfigFromJSON(configPath, "")
	fmt.Println(oc.Endpoint)

	assert.Nil(t, err)

}
