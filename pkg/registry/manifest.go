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

package registry

import (
	"log"
	"net/http"
	"sync"
)

type catalog struct {
	Repos []string `json:"repositories"`
}

type listTags struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

type manifest struct {
	contentType string
	blob        []byte
}

type manifests struct {
	// maps repo -> manifest tag/digest -> manifest
	manifests map[string]map[string]manifest
	lock      sync.RWMutex
	log       *log.Logger
}

func isManifest(req *http.Request) bool { _ = "STUB: not implemented"; return false }

func isTags(req *http.Request) bool { _ = "STUB: not implemented"; return false }

func isCatalog(req *http.Request) bool { _ = "STUB: not implemented"; return false }

// Returns whether this url should be handled by the referrers handler
func isReferrers(req *http.Request) bool { _ = "STUB: not implemented"; return false }

// https://github.com/opencontainers/distribution-spec/blob/master/spec.md#pulling-an-image-manifest
// https://github.com/opencontainers/distribution-spec/blob/master/spec.md#pushing-an-image
func (m *manifests) handle(resp http.ResponseWriter, req *http.Request) *regError {
	_ = "STUB: not implemented"
	return nil
}

// If the manifest is a manifest list, check that the manifest
// list's constituent manifests are already uploaded.
// This isn't strictly required by the registry API, but some
// registries require this.

// TODO: Probably want to do an existence check for blobs.

// Allow future references by target (tag) and immutable digest.
// See https://docs.docker.com/engine/reference/commandline/pull/#pull-an-image-by-digest-immutable-identifier.

func (m *manifests) handleTags(resp http.ResponseWriter, req *http.Request) *regError {
	_ = "STUB: not implemented"
	return nil
}

// https://github.com/opencontainers/distribution-spec/blob/b505e9cc53ec499edbd9c1be32298388921bb705/detail.md#tags-paginated
// Offset using last query parameter.

// Limit using n query parameter.

func (m *manifests) handleCatalog(resp http.ResponseWriter, req *http.Request) *regError {
	_ = "STUB: not implemented"
	return nil
}

// TODO: implement pagination

// TODO: implement handling of artifactType querystring
func (m *manifests) handleReferrers(resp http.ResponseWriter, req *http.Request) *regError {
	_ = "STUB: not implemented"
	// Ensure this is a GET request
	return nil
}

// Validate that incoming target is a valid digest

// At this point, we know the current digest references the target
