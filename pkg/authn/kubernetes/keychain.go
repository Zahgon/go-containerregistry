// Copyright 2022 Google LLC All Rights Reserved.
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

package kubernetes

import (
	"context"
	"net/url"

	"github.com/google/go-containerregistry/pkg/authn"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

const (
	// NoServiceAccount is a constant that can be passed via ServiceAccountName
	// to tell the keychain that looking up the service account is unnecessary.
	// This value cannot collide with an actual service account name because
	// service accounts do not allow spaces.
	NoServiceAccount = "no service account"
)

// Options holds configuration data for guiding credential resolution.
type Options struct {
	// Namespace holds the namespace inside of which we are resolving service
	// account and pull secret references to access the image.
	// If empty, "default" is assumed.
	Namespace string

	// ServiceAccountName holds the serviceaccount (within Namespace) as which a
	// Pod might access the image.  Service accounts may have image pull secrets
	// attached, so we lookup the service account to complete the keychain.
	// If empty, "default" is assumed.  To avoid a service account lookup, pass
	// NoServiceAccount explicitly.
	ServiceAccountName string

	// ImagePullSecrets holds the names of the Kubernetes secrets (scoped to
	// Namespace) containing credential data to use for the image pull.
	ImagePullSecrets []string

	// UseMountSecrets determines whether or not mount secrets in the ServiceAccount
	// should be considered. Mount secrets are those listed under the `.secrets`
	// attribute of the ServiceAccount resource. Ignored if ServiceAccountName is set
	// to NoServiceAccount.
	UseMountSecrets bool
}

// New returns a new authn.Keychain suitable for resolving image references as
// scoped by the provided Options.  It speaks to Kubernetes through the provided
// client interface.
func New(ctx context.Context, client kubernetes.Interface, opt Options) (authn.Keychain, error) {
	_ = "STUB: not implemented"
	return *new(authn.Keychain), nil
}

// Implement a Kubernetes-style authentication keychain.
// This needs to support roughly the following kinds of authentication:
//  1) The explicit authentication from imagePullSecrets on Pod
//  2) The semi-implicit authentication where imagePullSecrets are on the
//    Pod's service account.

// First, fetch all of the explicitly declared pull secrets

// Second, fetch all of the pull secrets attached to our service account,
// unless the user has explicitly specified that no service account lookup
// is desired.

// NewInCluster returns a new authn.Keychain suitable for resolving image references as
// scoped by the provided Options, constructing a kubernetes.Interface based on in-cluster
// authentication.
func NewInCluster(ctx context.Context, opt Options) (authn.Keychain, error) {
	_ = "STUB: not implemented"
	return *new(authn.Keychain), nil
}

type dockerConfigJSON struct {
	Auths map[string]authn.AuthConfig
}

// NewFromPullSecrets returns a new authn.Keychain suitable for resolving image references as
// scoped by the pull secrets.
func NewFromPullSecrets(ctx context.Context, secrets []corev1.Secret) (authn.Keychain, error) {
	_ = "STUB: not implemented"
	return *new(authn.Keychain), nil
}

// From: https://github.com/kubernetes/kubernetes/blob/0dcafb1f37ee522be3c045753623138e5b907001/pkg/credentialprovider/keyring.go

// The docker client allows exact matches:
//    foo.bar.com/namespace
// Or hostname matches:
//    foo.bar.com
// It also considers /v2/  and /v1/ equivalent to the hostname
// See ResolveAuthConfig in docker/registry/auth.go.

// We reverse sort in to give more specific (aka longer) keys priority
// when matching for creds

type keyring struct {
	index []string
	creds map[string][]authn.AuthConfig
}

func (keyring *keyring) Resolve(target authn.Resource) (authn.Authenticator, error) {
	_ = "STUB: not implemented"
	return *new(authn.Authenticator), nil
}

// both k and image are schemeless URLs because even though schemes are allowed
// in the credential configurations, we remove them when constructing the keyring

// urlsMatchStr is wrapper for URLsMatch, operating on strings instead of URLs.
func urlsMatchStr(glob string, target string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// parseSchemelessURL parses a schemeless url and returns a url.URL
// url.Parse require a scheme, but ours don't have schemes.  Adding a
// scheme to make url.Parse happy, then clear out the resulting scheme.
func parseSchemelessURL(schemelessURL string) (*url.URL, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// clear out the resulting scheme

// splitURL splits the host name into parts, as well as the port
func splitURL(url *url.URL) (parts []string, port string) {
	_ = "STUB: not implemented"
	return nil, ""
}

// could not parse port

// urlsMatch checks whether the given target url matches the glob url, which may have
// glob wild cards in the host name.
//
// Examples:
//
//	globURL=*.docker.io, targetURL=blah.docker.io => match
//	globURL=*.docker.io, targetURL=not.right.io   => no match
//
// Note that we don't support wildcards in ports and paths yet.
func urlsMatch(globURL *url.URL, targetURL *url.URL) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// port doesn't match

// host name does not have the same number of parts

// the path of the credential must be a prefix

// glob mismatch for some part

// everything matches

func toAuthenticator(configs []authn.AuthConfig) (authn.Authenticator, error) {
	_ = "STUB: not implemented"
	return *new(authn.Authenticator), nil
}
