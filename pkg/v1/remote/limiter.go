// Copyright 2026 Google LLC All Rights Reserved.
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
	"sync"
)

type pullLimiter struct {
	tokens chan struct{}
}

func newPullLimiter(jobs int) *pullLimiter { _ = "STUB: not implemented"; return nil }

func (l *pullLimiter) acquire(ctx context.Context) (func(), error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type limitedReadCloser struct {
	io.ReadCloser
	release func()
	once    sync.Once
}

func (l *limitedReadCloser) Close() error { _ = "STUB: not implemented"; return nil }
