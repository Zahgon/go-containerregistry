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
	"net/http"
	"net/url"
	"time"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
)

// Option is a functional option for List and Walk.
// TODO: Can we somehow reuse the remote options here?
type Option func(*lister) error

type lister struct {
	auth      authn.Authenticator
	transport http.RoundTripper
	repo      name.Repository
	client    *http.Client
	ctx       context.Context
	userAgent string
}

func newLister(repo name.Repository, options ...Option) (*lister, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// transport.Wrapper is a signal that consumers are opt-ing into providing their own transport without any additional wrapping.
// This is to allow consumers full control over the transports logic, such as providing retry logic.

// Wrap the transport in something that logs requests and responses.
// It's expensive to generate the dumps, so skip it if we're writing
// to nothing.

// Wrap the transport in something that can retry network flakes.

// Wrap this last to prevent transport.New from double-wrapping.

func (l *lister) list(repo name.Repository) (*Tags, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ECR returns an error if n > 1000:
// https://github.com/google/go-containerregistry/issues/681

// get responses until there is no next page

// We're dealing with GCR, just return directly.

// This isn't GCR, just append the tags and keep paginating.

// no next page

// getNextPageURL checks if there is a Link header in a http.Response which
// contains a link to the next page. If yes it returns the url.URL of the next
// page otherwise it returns nil.
func getNextPageURL(resp *http.Response) (*url.URL, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type rawManifestInfo struct {
	Size      string   `json:"imageSizeBytes"`
	MediaType string   `json:"mediaType"`
	Created   string   `json:"timeCreatedMs"`
	Uploaded  string   `json:"timeUploadedMs"`
	Tags      []string `json:"tag"`
}

// ManifestInfo is a Manifests entry is the output of List and Walk.
type ManifestInfo struct {
	Size      uint64    `json:"imageSizeBytes"`
	MediaType string    `json:"mediaType"`
	Created   time.Time `json:"timeCreatedMs"`
	Uploaded  time.Time `json:"timeUploadedMs"`
	Tags      []string  `json:"tag"`
}

func fromUnixMs(ms int64) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func toUnixMs(t time.Time) string { _ = "STUB: not implemented"; return "" }

// MarshalJSON implements json.Marshaler
func (m ManifestInfo) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON implements json.Unmarshaler
func (m *ManifestInfo) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// Tags is the result of List and Walk.
type Tags struct {
	Children  []string                `json:"child"`
	Manifests map[string]ManifestInfo `json:"manifest"`
	Name      string                  `json:"name"`
	Tags      []string                `json:"tags"`
}

// List calls /tags/list for the given repository.
func List(repo name.Repository, options ...Option) (*Tags, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WalkFunc is the type of the function called for each repository visited by
// Walk. This implements a similar API to filepath.Walk.
//
// The repo argument contains the argument to Walk as a prefix; that is, if Walk
// is called with "gcr.io/foo", which is a repository containing the repository
// "bar", the walk function will be called with argument "gcr.io/foo/bar".
// The tags and error arguments are the result of calling List on repo.
//
// TODO: Do we want a SkipDir error, as in filepath.WalkFunc?
type WalkFunc func(repo name.Repository, tags *Tags, err error) error

func walk(repo name.Repository, tags *Tags, walkFn WalkFunc, options ...Option) error {
	_ = "STUB: not implemented"

	// This shouldn't happen.
	return nil
}

// We don't expect this ever, so don't pass it through to walkFn.

// We made it!

// Walk recursively descends repositories, calling walkFn.
func Walk(root name.Repository, walkFn WalkFunc, options ...Option) error {
	_ = "STUB: not implemented"
	return nil
}
