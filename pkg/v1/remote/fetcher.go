// Copyright 2023 Google LLC All Rights Reserved.
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

package remote

import (
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/types"
)

const (
	kib           = 1024
	mib           = 1024 * kib
	manifestLimit = 100 * mib
)

// fetcher implements methods for reading from a registry.
type fetcher struct {
	target  resource
	client  *http.Client
	limiter *pullLimiter
}

func makeFetcher(ctx context.Context, target resource, o *options) (*fetcher, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// checkRedirectSSRF rejects HTTP redirects that cross from a public host to a
// private or link-local IP literal. This prevents a malicious registry from
// issuing a 302 to a cloud instance metadata service (e.g. 169.254.169.254)
// or another internal network address during blob or manifest downloads.
//
// Same-host redirects and redirects to non-IP hostnames (including DNS names
// that may resolve to private addresses) are allowed. The first redirect in
// the chain uses the original request URL as the "origin host" via
// req.Response.Request, falling back to req.URL when no prior response exists.
func checkRedirectSSRF(req *http.Request, via []*http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// same-host redirect is always allowed

func (f *fetcher) Do(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type resource interface {
	Scheme() string
	RegistryStr() string
	Scope(string) string

	authn.Resource
}

// url returns a url.Url for the specified path in the context of this remote image reference.
func (f *fetcher) url(resource, identifier string) url.URL {
	_ = "STUB: not implemented"
	return *new(url.URL)
}

// Default path if this is not a repository.

func (f *fetcher) get(ctx context.Context, ref name.Reference, acceptable []types.MediaType, platform v1.Platform) (*Descriptor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *fetcher) fetchManifest(ctx context.Context, ref name.Reference, acceptable []types.MediaType) ([]byte, *v1.Descriptor, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// If we can parse the digest from the header, and it's a signed schema 1
// manifest, let's use that for the digest to appease older registries.

// Validate the digest matches what we asked for, if pulling by digest.

// Failing to parse as a manifest should just be ignored.
// The manifest might not be valid, and that's okay.

// Per the OCI distribution spec, artifactType on the descriptor is
// set to the manifest's artifactType if present, otherwise it falls
// back to the config descriptor's mediaType.

// Do nothing for tags; I give up.
//
// We'd like to validate that the "Docker-Content-Digest" header matches what is returned by the registry,
// but so many registries implement this incorrectly that it's not worth checking.
//
// For reference:
// https://github.com/GoogleContainerTools/kaniko/issues/298

// Return all this info since we have to calculate it anyway.

func (f *fetcher) headManifest(ctx context.Context, ref name.Reference, acceptable []types.MediaType) (*v1.Descriptor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validate the digest matches what we asked for, if pulling by digest.

// Return all this info since we have to calculate it anyway.

func (f *fetcher) fetchBlob(ctx context.Context, size int64, h v1.Hash) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (f *fetcher) fetchBlobURL(ctx context.Context, u url.URL, size int64, h v1.Hash) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

// Do whatever we can.
// If we have an expected size and Content-Length doesn't match, return an error.
// If we don't have an expected size and we do have a Content-Length, use Content-Length.

func (f *fetcher) headBlob(ctx context.Context, h v1.Hash) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// validateForeignURL rejects foreign layer URLs that use a disallowed scheme
// or resolve to a private / link-local IP address (SSRF protection). DNS-based
// SSRF is out of scope, matching transport.validateRealmURL.
func validateForeignURL(rawURL string, insecure bool) error { _ = "STUB: not implemented"; return nil }

// fetchForeignBlobURL fetches a foreign-layer blob, validating every redirect
// destination through validateForeignURL (SSRF protection).
func (f *fetcher) fetchForeignBlobURL(ctx context.Context, u url.URL, size int64, h v1.Hash, insecure bool) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (f *fetcher) blobExists(ctx context.Context, h v1.Hash) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
