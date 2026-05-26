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

package crane

import (
	"errors"
)

// ErrRefusingToClobberExistingTag is returned when NoClobber is true and the
// tag already exists in the target registry/repo.
var ErrRefusingToClobberExistingTag = errors.New("refusing to clobber existing tag")

// Copy copies a remote image or index from src to dst.
func Copy(src, dst string, opt ...Option) error { _ = "STUB: not implemented"; return nil }

// If platform is explicitly set, don't copy the whole index, just the appropriate image.

// CopyRepository copies every tag from src to dst.
func CopyRepository(src, dst string, opt ...Option) error { _ = "STUB: not implemented"; return nil }

// TODO: It would be good to propagate noclobber down into remote so we can use Etags.

// Some registries create repository on first push, so listing tags will fail.
// If we see 404 or 403, assume we failed because the repository hasn't been created yet.
