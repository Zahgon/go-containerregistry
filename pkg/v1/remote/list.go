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

	"github.com/google/go-containerregistry/pkg/name"
)

// ListWithContext calls List with the given context.
//
// Deprecated: Use List and WithContext. This will be removed in a future release.
func ListWithContext(ctx context.Context, repo name.Repository, options ...Option) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// List calls /tags/list for the given repository, returning the list of tags
// in the "tags" property.
func List(repo name.Repository, options ...Option) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Tags struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
	Next string   `json:"next,omitempty"`
}

func (f *fetcher) listPage(ctx context.Context, repo name.Repository, next string, pageSize int) (*Tags, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getNextPageURL checks if there is a Link header in a http.Response which
// contains a link to the next page. If yes it returns the url.URL of the next
// page otherwise it returns nil.
func getNextPageURL(resp *http.Response) (*url.URL, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Lister struct {
	f        *fetcher
	repo     name.Repository
	pageSize int

	page *Tags
	err  error

	needMore bool
}

func (l *Lister) Next(ctx context.Context) (*Tags, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *Lister) HasNext() bool { _ = "STUB: not implemented"; return false }
