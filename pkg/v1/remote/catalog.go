// Copyright 2019 Google LLC All Rights Reserved.
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

	"github.com/google/go-containerregistry/pkg/name"
)

type Catalogs struct {
	Repos []string `json:"repositories"`
	Next  string   `json:"next,omitempty"`
}

// CatalogPage calls /_catalog, returning the list of repositories on the registry.
func CatalogPage(target name.Registry, last string, n int, options ...Option) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Catalog calls /_catalog, returning the list of repositories on the registry.
func Catalog(ctx context.Context, target name.Registry, options ...Option) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WithContext overrides the ctx passed directly.

func (f *fetcher) catalogPage(ctx context.Context, reg name.Registry, next string, pageSize int) (*Catalogs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Catalogger struct {
	f        *fetcher
	reg      name.Registry
	pageSize int

	page *Catalogs
	err  error

	needMore bool
}

func (l *Catalogger) Next(ctx context.Context) (*Catalogs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *Catalogger) HasNext() bool { _ = "STUB: not implemented"; return false }
