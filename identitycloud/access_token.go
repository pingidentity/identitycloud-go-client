// Copyright © 2025 Ping Identity Corporation

package identitycloud

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"
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

const maxOauthRetries = 4

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

	return oAuthExponentialBackOffRetry(func() (*oauth2.Token, error) {
		return s.internalToken(jwt)
	})
}

func (s ServiceAccountTokenSource) internalToken(jwt string) (*oauth2.Token, error) {
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
	if httpResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to request access token with service account. Status code: %d, body: %s", httpResp.StatusCode, body)
	}

	var token oauth2.Token
	err = json.Unmarshal(body, &token)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal access token response: %w, body: %s", err, body)
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

func oAuthExponentialBackOffRetry(f func() (*oauth2.Token, error)) (*oauth2.Token, error) {
	var token *oauth2.Token
	var err error
	backOffTime := time.Second
	var isRetryable bool

	for i := 0; i < maxOauthRetries; i++ {
		token, err = f()

		if err != nil {
			backOffTime, isRetryable = oAuthTestForRetryable(err, backOffTime)

			if isRetryable {
				log.Printf("Attempt %d failed: %v, backing off by %s.", i+1, err, backOffTime.String())
				time.Sleep(backOffTime)
				continue
			}
		}

		return token, err
	}

	log.Printf("Request failed after %d attempts", maxOauthRetries)

	return token, err // output the final error
}

func oAuthTestForRetryable(err error, currentBackoff time.Duration) (time.Duration, bool) {
	backoff := currentBackoff * 2

	retryAbleCodes := []int{
		429,
		500,
		502,
		503,
		504,
	}

	for _, v := range retryAbleCodes {
		if m, mErr := regexp.MatchString(fmt.Sprintf("%d ", v), err.Error()); mErr == nil && m {
			log.Printf("HTTP status code %d detected, available for retry", v)
			return backoff, true
		}
	}

	return backoff, false
}
