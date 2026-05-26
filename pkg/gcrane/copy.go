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

package gcrane

import (
	"context"

	"github.com/google/go-containerregistry/internal/retry"
	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/google"
)

// Keychain tries to use google-specific credential sources, falling back to
// the DefaultKeychain (config-file based).
var Keychain = authn.NewMultiKeychain(google.Keychain, authn.DefaultKeychain)

// GCRBackoff returns a retry.Backoff that is suitable for use with gcr.io.
//
// These numbers are based on GCR's posted quotas:
// https://cloud.google.com/container-registry/quotas
// - 50k requests per 10 minutes.
// -  1M requests per 24 hours.
//
// On error, we will wait for:
// - 6 seconds (in case of very short term 429s from GCS), then
// - 1 minute (in case of temporary network issues), then
// - 10 minutes (to get around GCR 10 minute quotas), then fail.
//
// TODO: In theory, we could keep retrying until the next day to get around the 1M limit.
func GCRBackoff() retry.Backoff { _ = "STUB: not implemented"; return *new(retry.Backoff) }

// Copy copies a remote image or index from src to dst.
func Copy(src, dst string, opts ...Option) error { _ = "STUB: not implemented"; return nil }

// Just reuse crane's copy logic with gcrane's credential logic.

// CopyRepository copies everything from the src GCR repository to the
// dst GCR repository.
func CopyRepository(ctx context.Context, src, dst string, opts ...Option) error {
	_ = "STUB: not implemented"
	return nil
}

type task struct {
	digest   string
	manifest google.ManifestInfo
	oldRepo  name.Repository
	newRepo  name.Repository
}

type copier struct {
	srcRepo name.Repository
	dstRepo name.Repository

	tasks chan task
	opt   *options
}

func newCopier(src, dst string, o *options) (*copier, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// A queue of size 2*jobs should keep each goroutine busy.

// recursiveCopy copies images from repo src to repo dst.
func recursiveCopy(ctx context.Context, src, dst string, o *options) error {
	_ = "STUB: not implemented"
	return nil
}

// If we hit an error when listing the repo, try re-listing with backoff.

// If we hit an error when trying to diff the repo, re-diff with backoff.

// Start walking the repo, enqueuing items in c.tasks.

// Pull items off of c.tasks and copy the images.

// If we hit an error when trying to copy the images,
// retry with backoff.

// copyRepo figures out the name for our destination repo (newRepo), lists the
// contents of newRepo, calculates the diff of what needs to be copied, then
// starts a goroutine to copy each image we need, and waits for them to finish.
func (c *copier) copyRepo(ctx context.Context, oldRepo name.Repository, tags *google.Tags) error {
	_ = "STUB: not implemented"
	return nil
}

// Figure out what we actually need to copy.

// This is a 404 code, so we just need to copy everything.

// Queue up every image as a task.

// copyImages starts a goroutine for each tag that points to the image
// oldRepo@digest, or just copies the image by digest if there are no tags.
func (c *copier) copyImages(_ context.Context, t task) error {
	_ = "STUB: not implemented"
	// We only have to explicitly copy by digest if there are no tags pointing to this manifest.
	return nil
}

// We only need to push the whole image once.

// If there's only one tag, we're done.

// Add the rest of the tags.

// Retry temporary errors, 429, and 500+ with backoff.
func backoffErrors(bo retry.Backoff, f func() error) error { _ = "STUB: not implemented"; return nil }

func hasStatusCode(err error, code int) bool { _ = "STUB: not implemented"; return false }

func isServerError(err error) bool { _ = "STUB: not implemented"; return false }

// rename figures out the name of the new repository to copy to, e.g.:
//
// $ gcrane cp -r gcr.io/foo gcr.io/baz
//
// rename("gcr.io/foo/bar") == "gcr.io/baz/bar"
func (c *copier) rename(repo name.Repository) (name.Repository, error) {
	_ = "STUB: not implemented"
	return *new(name.Repository), nil
}

// diffImages returns a map of digests to google.ManifestInfos for images or
// tags that are present in "want" but not in "have".
func diffImages(want, have map[string]google.ManifestInfo) map[string]google.ManifestInfo {
	_ = "STUB: not implemented"
	return nil
}

// Missing the whole image, we need to copy everything.

// Missing just some tags, add the ones we need to copy.

// subtractStringLists returns a list of strings that are in minuend and not
// in subtrahend; order is unimportant.
func subtractStringLists(minuend, subtrahend []string) []string {
	_ = "STUB: not implemented"
	return nil
}

func toStringSet(slice []string) map[string]struct{} { _ = "STUB: not implemented"; return nil }
