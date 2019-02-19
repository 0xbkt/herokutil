package herokutil

import (
	"context"
	"net/http"
	"os"

	"github.com/heroku/heroku-go/v3"
)

// NewClientWithToken returns a new HTTP client with Bearer token `token`
// for use with Heroku.
func NewClientWithToken(token string) *http.Client {
	return &http.Client{
		Transport: &heroku.Transport{
			BearerToken: token,
		},
	}
}

// LoadEnvOf loads the config variables of appID into current environment.
func LoadEnvOf(svc *heroku.Service, appID string) (err error) {
	cfg, err := svc.ConfigVarInfoForApp(context.Background(), appID)
	if err != nil {
		return
	}

	for k, v := range cfg {
		os.Setenv(k, *v)
	}

	return
}

