package config

import (
	"context"
	"log"

	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

var (
	OIDCProvider *oidc.Provider
	OAuth2Config oauth2.Config
	Verifier     *oidc.IDTokenVerifier
)

func InitOIDC() {
	issuer := "https://accounts.google.com"
	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	redirectURL := "http://localhost:8080/callback"

	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, issuer)
	if err != nil {
		log.Fatalf("Failed to get OIDC provider: %v", err)
	}

	OIDCProvider = provider
	OAuth2Config = oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  redirectURL,
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	Verifier = provider.Verifier(&oidc.Config{ClientID: clientID})
}
