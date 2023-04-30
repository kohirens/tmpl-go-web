package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/b01/{{ .codeName }}/internal/cli"
	"github.com/kohirens/stdlib"
	"golang.org/x/oauth2/clientcredentials"
	"net/http"
	"os"
)

// GetClient Get an HTTP Client configured for OAuth2. authFile is a JSON File
// containing any of the OAuth2 Config 2-legged OAuth2 flow, with both the
// client application information and the server's endpoint URLs., for example:
//
//	ClientID       string
//	ClientSecret   string
//	TokenURL       string
//	Scopes         []string
//	EndpointParams url.Values
//	AuthStyle      oauth2.AuthStyle
func GetClient(authFile string) (*http.Client, error) {
	authConf := &clientcredentials.Config{}

	if !stdlib.PathExist(authFile) {
		return nil, fmt.Errorf(cli.Stderr.NoAuthFile)
	}

	authBytes, err1 := os.ReadFile(authFile)
	if err1 != nil {
		return nil, err1
	}

	if e := json.Unmarshal(authBytes, authConf); e != nil {
		return nil, e
	}

	// Get an HTTP client
	return authConf.Client(context.Background()), nil
}
