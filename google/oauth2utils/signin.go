package oauth2utils

import (
	"io/ioutil"
	"os"

	"github.com/yongtin/goutils/files"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// OAuth2ConfigFromJSON return an oauth client from config json file.
func OAuth2ConfigFromJSON(ConfigJSONFile string, scope string) (o *oauth2.Config, err error) {
	if _, err := files.IsFile(ConfigJSONFile); err != nil {
		return o, err
	}
	fp, err := os.Open(ConfigJSONFile)
	if err != nil {
		return o, err
	}
	configData, err := ioutil.ReadAll(fp)
	if err != nil {
		return o, err
	}

	return google.ConfigFromJSON(configData, scope)
}
