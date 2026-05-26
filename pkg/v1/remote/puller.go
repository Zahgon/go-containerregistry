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
	"github.com/google/go-containerregistry/pkg/v1/types"
)

type Puller struct {
	o *options

	// map[resource]*reader
	readers sync.Map
}

func NewPuller(options ...Option) (*Puller, error) { _ = "STUB: not implemented"; return nil, nil }

func newPuller(o *options) *Puller { _ = "STUB: not implemented"; return nil }

type reader struct {
	// in
	target resource
	o      *options

	// f()
	once sync.Once

	// out
	f   *fetcher
	err error
}

// this will run once per reader instance
func (r *reader) init(ctx context.Context) error {
	r.once.Do(func() {
		r.f, r.err = makeFetcher(ctx, r.target, r.o)
	})
	return r.err
}

func (p *Puller) fetcher(ctx context.Context, target resource) (*fetcher, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Head is like remote.Head, but avoids re-authenticating when possible.
func (p *Puller) Head(ctx context.Context, ref name.Reference) (*v1.Descriptor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get is like remote.Get, but avoids re-authenticating when possible.
func (p *Puller) Get(ctx context.Context, ref name.Reference) (*Descriptor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Puller) get(ctx context.Context, ref name.Reference, acceptable []types.MediaType, platform v1.Platform) (*Descriptor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Layer is like remote.Layer, but avoids re-authenticating when possible.
func (p *Puller) Layer(ctx context.Context, ref name.Digest) (v1.Layer, error) {
	_ = "STUB: not implemented"
	return *new(v1.Layer), nil
}

// List lists tags in a repo and handles pagination, returning the full list of tags.
func (p *Puller) List(ctx context.Context, repo name.Repository) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Lister lists tags in a repo and returns a Lister for paginating through the results.
func (p *Puller) Lister(ctx context.Context, repo name.Repository) (*Lister, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Puller) lister(ctx context.Context, repo name.Repository, pageSize int) (*Lister, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Catalog lists repos in a registry and handles pagination, returning the full list of repos.
func (p *Puller) Catalog(ctx context.Context, reg name.Registry) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Puller) catalog(ctx context.Context, reg name.Registry, pageSize int) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Catalogger lists repos in a registry and returns a Catalogger for paginating through the results.
func (p *Puller) Catalogger(ctx context.Context, reg name.Registry) (*Catalogger, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Puller) catalogger(ctx context.Context, reg name.Registry, pageSize int) (*Catalogger, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Puller) referrers(ctx context.Context, d name.Digest, filter map[string]string) (v1.ImageIndex, error) {
	_ = "STUB: not implemented"
	return *new(v1.ImageIndex), nil
}
