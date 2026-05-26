// Copyright 2018 Google LLC All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package google

import (
	"context"
	"os/exec"

	"github.com/google/go-containerregistry/pkg/authn"
	"golang.org/x/oauth2"
)

const cloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

// gcloudBin is replaced in tests to mock detecting the gcloud binary.
var gcloudBin = "gcloud"

// getGcloudCmd is replaced in tests to drive the gcloud mock.
var getGcloudCmd = func(ctx context.Context) *exec.Cmd {
	// This is odd, but basically what docker-credential-gcr does.
	//
	// config-helper is undocumented, but it's purportedly the only supported way
	// of accessing tokens (`gcloud auth print-access-token` is discouraged).
	//
	// --force-auth-refresh means we are getting a token that is valid for about
	// an hour (we reuse it until it's expired).
	return exec.CommandContext(ctx, gcloudBin, "config", "config-helper", "--force-auth-refresh", "--format=json(credential)")
}

// NewEnvAuthenticator returns an authn.Authenticator that generates access
// tokens from the environment we're running in.
//
// See: https://godoc.org/golang.org/x/oauth2/google#FindDefaultCredentials
func NewEnvAuthenticator(ctx context.Context) (authn.Authenticator, error) {
	_ = "STUB: not implemented"
	return *new(authn.Authenticator), nil
}

// NewGcloudAuthenticator returns an oauth2.TokenSource that generates access
// tokens by shelling out to the gcloud sdk.
func NewGcloudAuthenticator(ctx context.Context) (authn.Authenticator, error) {
	_ = "STUB: not implemented"
	return *new(authn.Authenticator), nil
}

// gcloud is not available, fall back to anonymous

// Attempt to fetch a token to ensure gcloud is installed and we can run it.

// NewJSONKeyAuthenticator returns a Basic authenticator which uses Service Account
// as a way of authenticating with Google Container Registry.
// More information: https://cloud.google.com/container-registry/docs/advanced-authentication#json_key_file
func NewJSONKeyAuthenticator(serviceAccountJSON string) authn.Authenticator {
	_ = "STUB: not implemented"
	return *new(authn.Authenticator)
}

// NewTokenAuthenticator returns an oauth2.TokenSource that generates access
// tokens by using the Google SDK to produce JWT tokens from a Service Account.
// More information: https://godoc.org/golang.org/x/oauth2/google#JWTAccessTokenSourceFromJSON
func NewTokenAuthenticator(serviceAccountJSON string, scope string) (authn.Authenticator, error) {
	_ = "STUB: not implemented"
	return *new(authn.Authenticator), nil
}

// NewTokenSourceAuthenticator converts an oauth2.TokenSource into an authn.Authenticator.
func NewTokenSourceAuthenticator(ts oauth2.TokenSource) authn.Authenticator {
	_ = "STUB: not implemented"
	return *new(authn.Authenticator)
}

// tokenSourceAuth turns an oauth2.TokenSource into an authn.Authenticator.
type tokenSourceAuth struct {
	oauth2.TokenSource
}

// Authorization implements authn.Authenticator.
func (tsa *tokenSourceAuth) Authorization() (*authn.AuthConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// gcloudOutput represents the output of the gcloud command we invoke.
//
// `gcloud config config-helper --format=json(credential)` looks something like:
//
//	{
//	  "credential": {
//	    "access_token": "supersecretaccesstoken",
//	    "token_expiry": "2018-12-02T04:08:13Z"
//	  }
//	}
type gcloudOutput struct {
	Credential struct {
		AccessToken string `json:"access_token"`
		TokenExpiry string `json:"token_expiry"`
	} `json:"credential"`
}

type gcloudSource struct {
	ctx context.Context

	// This is passed in so that we mock out gcloud and test Token.
	exec func(ctx context.Context) *exec.Cmd
}

// Token implements oauath2.TokenSource.
func (gs gcloudSource) Token() (*oauth2.Token, error) { _ = "STUB: not implemented"; return nil, nil }

// Don't attempt to interpret stderr, just pass it through.
