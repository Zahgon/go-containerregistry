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
	"sync"

	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/partial"
	"github.com/google/go-containerregistry/pkg/v1/types"
	"golang.org/x/sync/errgroup"
)

type manifest interface {
	Taggable
	partial.Describable
}

// key is either v1.Hash or v1.Layer (for stream.Layer)
type workers struct {
	// map[v1.Hash|v1.Layer]*sync.Once
	onces sync.Map

	// map[v1.Hash|v1.Layer]error
	errors sync.Map
}

func nop() error { _ = "STUB: not implemented"; return nil }

func (w *workers) err(digest v1.Hash) error { _ = "STUB: not implemented"; return nil }

func (w *workers) Do(digest v1.Hash, f func() error) error {
	_ = "STUB: not implemented"
	// We don't care if it was loaded or not because the sync.Once will do it for us.
	return nil
}

// Allow this to be retried by another caller.

func (w *workers) Stream(layer v1.Layer, f func() error) error {
	_ = "STUB: not implemented"
	// We don't care if it was loaded or not because the sync.Once will do it for us.
	return nil
}

type Pusher struct {
	o *options

	// map[name.Repository]*repoWriter
	writers sync.Map
}

func NewPusher(options ...Option) (*Pusher, error) { _ = "STUB: not implemented"; return nil, nil }

func newPusher(o *options) *Pusher { _ = "STUB: not implemented"; return nil }

func (p *Pusher) writer(ctx context.Context, repo name.Repository, o *options) (*repoWriter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Pusher) Put(ctx context.Context, ref name.Reference, t Taggable) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Pusher) Push(ctx context.Context, ref name.Reference, t Taggable) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Pusher) Upload(ctx context.Context, repo name.Repository, l v1.Layer) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Pusher) Delete(ctx context.Context, ref name.Reference) error {
	_ = "STUB: not implemented"
	// Use a transport scoped for delete. Requesting DeleteScope (which
	// includes the "delete" action) allows registries that require an
	// explicit delete permission—such as IBM Cloud Container Registry—to
	// grant access.
	return nil
}

// TODO(jason): If the manifest had a `subject`, and if the registry
// doesn't support Referrers, update the index pointed to by the
// subject's fallback tag to remove the descriptor for this manifest.

type repoWriter struct {
	repo name.Repository
	o    *options
	once sync.Once

	w   *writer
	err error

	work *workers
}

// this will run once per repoWriter instance
func (rw *repoWriter) init(ctx context.Context) error {
	rw.once.Do(func() {
		rw.work = &workers{}
		rw.w, rw.err = makeWriter(ctx, rw.repo, nil, rw.o)
	})
	return rw.err
}

func (rw *repoWriter) writeDeps(ctx context.Context, m manifest) error {
	_ = "STUB: not implemented"
	return nil
}

// This has no deps, not an error (e.g. something you want to just PUT).

type describable struct {
	desc v1.Descriptor
}

func (d describable) Digest() (v1.Hash, error) {
	_ = "STUB: not implemented"
	return *new(v1.Hash), nil
}

func (d describable) Size() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (d describable) MediaType() (types.MediaType, error) {
	_ = "STUB: not implemented"
	return *new(types.MediaType), nil
}

type tagManifest struct {
	Taggable
	partial.Describable
}

func taggableToManifest(t Taggable) (manifest, error) {
	_ = "STUB: not implemented"
	return *new(manifest), nil
}

// A reasonable default if Taggable doesn't implement MediaType.

func (rw *repoWriter) writeManifest(ctx context.Context, ref name.Reference, t Taggable) error {
	_ = "STUB: not implemented"
	return nil
}

// This may be a lazy child where we have no ref until digest is computed.

// For tags, we want to do this check outside of our Work.Do closure because
// we don't want to dedupe based on the manifest digest.

// The following work.Do will get deduped by digest, so it won't happen unless
// this tag happens to be the first commitManifest to run for that digest.

// Only runs for tags that got deduped by digest.

func (rw *repoWriter) writeChildren(ctx context.Context, idx v1.ImageIndex) error {
	_ = "STUB: not implemented"
	return nil
}

func (rw *repoWriter) writeChild(ctx context.Context, child partial.Describable, g *errgroup.Group) error {
	_ = "STUB: not implemented"
	return nil
}

// For recursive index, we want to do a depth-first launching of goroutines
// to avoid deadlocking.
//
// Note that this is rare, so the impact of this should be really small.

// This can't happen.

// TODO: Consider caching some representation of the tags/digests in the destination
// repository as a hint to avoid this optimistic check in cases where we will most
// likely have to do a PUT anyway, e.g. if we are overwriting a tag we just wrote.
func (rw *repoWriter) manifestExists(ctx context.Context, ref name.Reference, t Taggable) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Possibly due to streaming layers.

// We treat a 403 here as non-fatal because this existence check is an optimization and
// some registries will return a 403 instead of a 404 in certain situations.
// E.g. https://jfrog.atlassian.net/browse/RTFACT-13797

// Mark that we saw this digest in the registry so we don't have to check it again.

func (rw *repoWriter) commitManifest(ctx context.Context, ref name.Reference, m manifest) error {
	_ = "STUB: not implemented"
	return nil
}

func (rw *repoWriter) writeLayers(pctx context.Context, img v1.Image) error {
	_ = "STUB: not implemented"
	return nil
}

func (rw *repoWriter) writeLayer(ctx context.Context, l v1.Layer) error {
	_ = "STUB: not implemented"
	// Skip any non-distributable things.
	return nil
}

func (rw *repoWriter) lazyWriteLayer(ctx context.Context, l v1.Layer) error {
	_ = "STUB: not implemented"
	return nil
}

// Mark this upload completed.
