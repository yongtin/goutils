package oauth2

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"

	"github.com/pkg/browser"
	"github.com/yongtin/goutils/files"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// GetoAuth2ConfigFromJSON return an oAuth2 client from config json file
func GetoAuth2ConfigFromJSON(ConfigJSONFile string, scope string) (o *oauth2.Config, err error) {
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

// GetHTTPClient return an oAuth2 validated http client
func GetHTTPClient(config *oauth2.Config) (httpClient *http.Client, err error) {
	// Use cache oAuth2 token (by scope and time) when possible
	// If cache has expired:
	//   use redirect_url: create a temp webserver to catch oAuth2 validation code via redirect_url
	//   use the validation code to retrieve token and save locally as cache
	// reference code: https://developers.google.com/people/quickstart/go

	tok, err := getTokenFromFile("token.json")
	if err != nil {
		tok, err = getTokenFromWeb(config)
		if err != nil {
			return httpClient, err
		}
		_ = saveTokenToFile("token.json", tok)
	}

	return config.Client(context.Background(), tok), nil
}

// getTokenFromFile retrieve previously saved oauth2 json token
func getTokenFromFile(filepath string) (*oauth2.Token, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	token := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(token)
	return token, err
}

// saveTokenToFile saves oAuth code.Exchanged token as json cache
func saveTokenToFile(filepath string, token *oauth2.Token) (err error) {
	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	defer f.Close()
	json.NewEncoder(f).Encode(token)
	return err
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) (*oauth2.Token, error) {
	// We will run two goroutines:
	// - a goroutine for starting and catching code
	// - and a main (blocking) goroutine to open browser to trigger sending code
	//
	// Note:
	// using access_type=offline to retrieve refresh token for continuous access
	//
	// reference: https://developers.google.com/identity/protocols/oauth2/openid-connect
	//
	config.RedirectURL = "http://localhost:63333/"
	u, _ := url.Parse(config.RedirectURL)

	var chAuthCode = make(chan string) // channel for authCode communication
	var wg sync.WaitGroup

	fmt.Printf("Launching local server %s to catch authCode ...\n", config.RedirectURL)
	wg.Add(1) // adding one waitgroup below
	go func() {
		defer wg.Done() // signal a wait() call when we are done
		http.HandleFunc(u.Path, func(w http.ResponseWriter, r *http.Request) {
			var q = r.URL.Query()
			if q.Get("state") == "state-get-token-from-web" {
				chAuthCode <- q.Get("code")
			} else {
				chAuthCode <- "invalid"
			}

		})
		var serverPort = strings.Split(u.Host, ":")[1]
		http.ListenAndServe(fmt.Sprintf(":%s", serverPort), nil)
	}()

	authURL := config.AuthCodeURL("state-get-token-from-web", oauth2.AccessTypeOffline)

	fmt.Printf("Launching browser at URL: %s\n", authURL)
	browser.OpenURL(authURL)

	var authCode = <-chAuthCode
	fmt.Printf("Received authCode: %s\n", authCode)
	return config.Exchange(context.TODO(), authCode)
}
