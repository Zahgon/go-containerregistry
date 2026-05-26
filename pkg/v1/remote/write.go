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

package remote

import (
	"context"
	"net/http"
	"net/url"
	"sync"

	"github.com/google/go-containerregistry/internal/retry"
	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/types"
)

// Taggable is an interface that enables a manifest PUT (e.g. for tagging).
type Taggable interface {
	RawManifest() ([]byte, error)
}

// Write pushes the provided img to the specified image reference.
func Write(ref name.Reference, img v1.Image, options ...Option) (rerr error) {
	_ = "STUB: not implemented"
	return nil
}

// writer writes the elements of an image to a remote image reference.
type writer struct {
	repo      name.Repository
	auth      authn.Authenticator
	transport http.RoundTripper

	client *http.Client

	progress  *progress
	backoff   Backoff
	predicate retry.Predicate

	scopeLock sync.Mutex
	// Keep track of scopes that we have already requested.
	scopeSet map[string]struct{}
	scopes   []string
}

// makeDeleteClient returns an HTTP client whose token includes the "delete"
// action so that registries requiring an explicit delete permission grant
// access for manifest deletion.
func makeDeleteClient(ctx context.Context, repo name.Repository, o *options) (*http.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func makeWriter(ctx context.Context, repo name.Repository, ls []v1.Layer, o *options) (*writer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// url returns a url.Url for the specified path in the context of this remote image reference.
func (w *writer) url(path string) url.URL { _ = "STUB: not implemented"; return *new(url.URL) }

func (w *writer) maybeUpdateScopes(ctx context.Context, ml *MountableLayer) error {
	_ = "STUB: not implemented"
	return nil
}

// nextLocation extracts the fully-qualified URL to which we should send the next request in an upload sequence.
func (w *writer) nextLocation(resp *http.Response) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// If the location header returned is just a url path, then fully qualify it.
// We cannot simply call w.url, since there might be an embedded query string.

// Reject Location headers that redirect to a DIFFERENT host that resolves to
// a private or link-local IP literal. A malicious or compromised registry can
// respond to a blob upload initiation (POST /v2/.../blobs/uploads/) with a
// crafted Location header pointing at an internal service, causing the client
// to send subsequent PATCH/PUT requests (including the layer data as the body)
// to that internal address. Pre-signed blob URLs from cloud storage providers
// (GCS, S3, Azure Blob) use public hostnames, so legitimate cross-host
// redirects are unaffected.
//
// Same-host redirects (e.g. a different path on the registry itself) are
// always allowed regardless of whether the registry IP is private.

// checkExistingBlob checks if a blob exists already in the repository by making a
// HEAD request to the blob store API.  GCR performs an existence check on the
// initiation if "mount" is specified, even if no "from" sources are specified.
// However, this is not broadly applicable to all registries, e.g. ECR.
func (w *writer) checkExistingBlob(ctx context.Context, h v1.Hash) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// initiateUpload initiates the blob upload, which starts with a POST that can
// optionally include the hash of the layer and a list of repositories from
// which that layer might be read. On failure, an error is returned.
// On success, the layer was either mounted (nothing more to do) or a blob
// upload was initiated and the body of that blob should be sent to the returned
// location.
func (w *writer) initiateUpload(ctx context.Context, from, mount, origin string) (location string, mounted bool, err error) {
	_ = "STUB: not implemented"
	return "", false, nil
}

// Quay will fail if we specify a "mount" without a "from".

// Make the request to initiate the blob upload.

// https://github.com/google/go-containerregistry/issues/1679

// https://github.com/google/go-containerregistry/issues/1404

// Check the response code to determine the result.

// We're done, we were able to fast-path.

// Proceed to PATCH, upload has begun.

// streamBlob streams the contents of the blob to the specified location.
// On failure, this will return an error.  On success, this will return the location
// header indicating how to commit the streamed blob.
func (w *writer) streamBlob(ctx context.Context, layer v1.Layer, streamLocation string) (commitLocation string, rerr error) {
	_ = "STUB: not implemented"
	return "", nil
}

// We can't retry streaming layers.

// If we know the size, set it.

// The blob has been uploaded, return the location header indicating
// how to commit this layer.

// commitBlob commits this blob by sending a PUT to the location returned from
// streaming the blob.
func (w *writer) commitBlob(ctx context.Context, location, digest string) error {
	_ = "STUB: not implemented"
	return nil
}

// incrProgress increments and sends a progress update, if WithProgress is used.
func (w *writer) incrProgress(written int64) { _ = "STUB: not implemented"; return }

// uploadOne performs a complete upload of a single layer.
func (w *writer) uploadOne(ctx context.Context, l v1.Layer) error {
	_ = "STUB: not implemented"
	return nil
}

// If we know the digest, this isn't a streaming layer. Do an existence
// check so we can skip uploading the layer if possible.

// This keeps breaking with DockerHub.
// https://github.com/google/go-containerregistry/issues/1741

// Only log layers with +json or +yaml. We can let through other stuff if it becomes popular.
// TODO(opencontainers/image-spec#791): Would be great to have an actual parser.

type withMediaType interface {
	MediaType() (types.MediaType, error)
}

// This is really silly, but go interfaces don't let me satisfy remote.Taggable
// with remote.Descriptor because of name collisions between method names and
// struct fields.
//
// Use reflection to either pull the v1.Descriptor out of remote.Descriptor or
// create a descriptor based on the RawManifest and (optionally) MediaType.
func unpackTaggable(t Taggable) ([]byte, *v1.Descriptor, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// A reasonable default if Taggable doesn't implement MediaType.

// commitSubjectReferrers is responsible for updating the fallback tag manifest to track descriptors referring to a subject for registries that don't yet support the Referrers API.
// TODO: use conditional requests to avoid race conditions
func (w *writer) commitSubjectReferrers(ctx context.Context, sub name.Digest, add v1.Descriptor) error {
	_ = "STUB: not implemented"
	// Check if the registry supports Referrers API.
	// TODO: This should be done once per registry, not once per subject.
	return nil
}

// The registry supports Referrers API. The registry is responsible for updating the referrers list.

// The registry doesn't support Referrers API, we need to update the manifest tagged with the fallback tag.
// Make the request to GET the current manifest.

// Not found just means there are no attachments. Start with an empty index.

// The digest is already attached, nothing to do.

// Append the new descriptor to the index.

// Sort the manifests for reproducibility.

type fallbackTaggable struct {
	im v1.IndexManifest
}

func (f fallbackTaggable) RawManifest() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
func (f fallbackTaggable) MediaType() (types.MediaType, error) {
	_ = "STUB: not implemented"
	return *new(types.MediaType), nil
}

// commitManifest does a PUT of the image's manifest.
func (w *writer) commitManifest(ctx context.Context, t Taggable, ref name.Reference) error {
	_ = "STUB: not implemented"
	// If the manifest refers to a subject, we need to check whether we need to update the fallback tag manifest.
	return nil
}

// Make the request to PUT the serialized manifest

// If the manifest referred to a subject, we may need to update the fallback tag manifest.
// TODO: If this fails, we'll retry the whole upload. We should retry just this part.

// The image was successfully pushed!

func scopesForUploadingImage(repo name.Repository, layers []v1.Layer) []string {
	_ = "STUB: not implemented"
	// use a map as set to remove duplicates scope strings
	return nil
}

// we will add push scope for ref.Context() after the loop.
// for now we ask pull scope for references of the same registry

// Push scope should be the first element because a few registries just look at the first scope to determine access.

// WriteIndex pushes the provided ImageIndex to the specified image reference.
// WriteIndex will attempt to push all of the referenced manifests before
// attempting to push the ImageIndex, to retain referential integrity.
func WriteIndex(ref name.Reference, ii v1.ImageIndex, options ...Option) (rerr error) {
	_ = "STUB: not implemented"
	return nil
}

// WriteLayer uploads the provided Layer to the specified repo.
func WriteLayer(repo name.Repository, layer v1.Layer, options ...Option) (rerr error) {
	_ = "STUB: not implemented"
	return nil
}

// Tag adds a tag to the given Taggable via PUT /v2/.../manifests/<tag>
//
// Notable implementations of Taggable are v1.Image, v1.ImageIndex, and
// remote.Descriptor.
//
// If t implements MediaType, we will use that for the Content-Type, otherwise
// we will default to types.DockerManifestSchema2.
//
// Tag does not attempt to write anything other than the manifest, so callers
// should ensure that all blobs or manifests that are referenced by t exist
// in the target registry.
func Tag(tag name.Tag, t Taggable, options ...Option) error { _ = "STUB: not implemented"; return nil }

// Put adds a manifest from the given Taggable via PUT /v1/.../manifest/<ref>
//
// Notable implementations of Taggable are v1.Image, v1.ImageIndex, and
// remote.Descriptor.
//
// If t implements MediaType, we will use that for the Content-Type, otherwise
// we will default to types.DockerManifestSchema2.
//
// Put does not attempt to write anything other than the manifest, so callers
// should ensure that all blobs or manifests that are referenced by t exist
// in the target registry.
func Put(ref name.Reference, t Taggable, options ...Option) error {
	_ = "STUB: not implemented"
	return nil
}

// Push uploads the given Taggable to the specified reference.
func Push(ref name.Reference, t Taggable, options ...Option) (rerr error) {
	_ = "STUB: not implemented"
	return nil
}
