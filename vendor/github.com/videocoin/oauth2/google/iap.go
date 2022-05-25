// Package google provides an oauth2 token source for authenticating with
// Google Identity Aware Proxy.
package google

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"cloud.google.com/go/compute/metadata"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"golang.org/x/xerrors"
)

const envVar = "GOOGLE_APPLICATION_CREDENTIALS"

func IAPTokenSource(ctx context.Context, oauthClientID string, filename string) (oauth2.TokenSource, error) {
	return getTokenSource(ctx, filename, oauthClientID)
}

func getTokenSource(ctx context.Context, filename string, audience string) (oauth2.TokenSource, error) {
	if filename == "" {
		if f := os.Getenv(envVar); f != "" {
			filename = f
		}
	}

	if filename == "" {
		f := wellKnownFile()
		if _, err := os.Stat(f); err != nil {
			if !os.IsNotExist(err) {
				return nil, err
			}
		} else {
			filename = f
		}
	}

	if filename != "" {
		cfg, err := readCredentialsFile(filename)
		if err != nil {
			return nil, err
		}
		cfg.UseIDToken = true
		cfg.PrivateClaims = map[string]interface{}{
			"target_audience": audience,
		}
		return cfg.TokenSource(ctx), nil
	}

	if metadata.OnGCE() {
		return MetadataTokenSource(audience), nil
	}

	return nil, errors.New("unable to determine credentials source")
}

func wellKnownFile() string {
	const f = "application_default_credentials.json"
	if runtime.GOOS == "windows" {
		return filepath.Join(os.Getenv("APPDATA"), "gcloud", f)
	}
	return filepath.Join(os.Getenv("HOME"), ".config", "gcloud", f)
}

func readCredentialsFile(filename string) (*jwt.Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	c, err := google.JWTConfigFromJSON(data)
	if err != nil {
		return nil, xerrors.Errorf("failed to read service account from %s %w", filename, err)
	}
	return c, nil
}

func MetadataTokenSource(audience string) oauth2.TokenSource {
	return oauth2.ReuseTokenSource(nil, &metadataSource{audience: audience})
}

type metadataSource struct {
	audience string
}

// see https://cloud.google.com/run/docs/authenticating/service-to-service
func (m *metadataSource) Token() (*oauth2.Token, error) {
	data, err := metadata.Get("instance/service-accounts/default/identity?audience=" + m.audience)
	if err != nil {
		return nil, xerrors.Errorf("failed to get token from metadata service: %w", err)
	}

	return &oauth2.Token{
		AccessToken: data,
		TokenType:   "Bearer",
		Expiry:      time.Now().Add(time.Minute * 30),
	}, nil
}
