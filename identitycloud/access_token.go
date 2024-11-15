package identitycloud

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"golang.org/x/oauth2"
)

var (
	_ oauth2.TokenSource = ServiceAccountTokenSource{}
)

type ServiceAccountTokenSource struct {
	TenantFqdn               string
	ServiceAccountId         string
	ServiceAccountPrivateKey string
	Scopes                   []string
}

func (s ServiceAccountTokenSource) Token() (*oauth2.Token, error) {
	jwt, err := s.buildJwt()
	if err != nil {
		return nil, fmt.Errorf("failed to build jwt: %w", err)
	}

	formData := url.Values{}
	formData.Add("client_id", "service-account")
	formData.Add("grant_type", "urn:ietf:params:oauth:grant-type:jwt-bearer")
	formData.Add("assertion", jwt)
	formData.Add("scope", strings.Join(s.Scopes, " "))

	httpClient := &http.Client{}
	httpResp, err := httpClient.Post(s.tokenUrl(), "application/x-www-form-urlencoded", strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to request access token with service account: %w", err)
	}

	defer httpResp.Body.Close()
	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read access token response body: %w", err)
	}

	var token oauth2.Token
	err = json.Unmarshal(body, &token)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal access token response: %w", err)
	}

	token.Expiry = time.Now().Add(time.Duration(token.ExpiresIn) * time.Second)

	return &token, nil
}

func (s ServiceAccountTokenSource) buildJwt() (string, error) {
	key, err := jwk.ParseKey([]byte(s.ServiceAccountPrivateKey))
	if err != nil {
		return "", fmt.Errorf("failed to parse private key: %w", err)
	}

	tok, err := jwt.NewBuilder().
		Issuer(s.ServiceAccountId).
		Subject(s.ServiceAccountId).
		Expiration(time.Now().Add(time.Second * 899)).
		JwtID(uuid.New().String()).
		Audience([]string{s.tokenUrl()}).
		Build()
	if err != nil {
		return "", fmt.Errorf("failed to build jwt: %w", err)
	}

	token, err := jwt.Sign(tok, jwt.WithKey(jwa.RS256, key))
	if err != nil {
		return "", fmt.Errorf("failed to sign jwt: %w", err)
	}

	return string(token), nil
}

func (s ServiceAccountTokenSource) tokenUrl() string {
	return fmt.Sprintf("https://%s:443/am/oauth2/access_token", s.TenantFqdn)
}
