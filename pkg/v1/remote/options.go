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
	"errors"
	"io"
	"net"
	"net/http"
	"syscall"
	"time"

	"github.com/google/go-containerregistry/internal/retry"
	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/logs"
	v1 "github.com/google/go-containerregistry/pkg/v1"
)

// Option is a functional option for remote operations.
type Option func(*options) error

type options struct {
	auth                           authn.Authenticator
	keychain                       authn.Keychain
	transport                      http.RoundTripper
	context                        context.Context
	jobs                           int
	userAgent                      string
	allowNondistributableArtifacts bool
	progress                       *progress
	retryBackoff                   Backoff
	retryPredicate                 retry.Predicate
	retryStatusCodes               []int
	limiter                        *pullLimiter

	// Only these options can overwrite Reuse()d options.
	platform v1.Platform
	pageSize int
	filter   map[string]string

	// Set by Reuse, we currently store one or the other.
	puller *Puller
	pusher *Pusher
}

var defaultPlatform = v1.Platform{
	Architecture: "amd64",
	OS:           "linux",
}

// Backoff is an alias of retry.Backoff to expose this configuration option to consumers of this lib
type Backoff = retry.Backoff

var defaultRetryPredicate retry.Predicate = func(err error) bool {
	// Various failure modes here, as we're often reading from and writing to
	// the network.
	if retry.IsTemporary(err) || errors.Is(err, io.ErrUnexpectedEOF) || errors.Is(err, io.EOF) || errors.Is(err, syscall.EPIPE) || errors.Is(err, syscall.ECONNRESET) || errors.Is(err, net.ErrClosed) {
		logs.Warn.Printf("retrying %v", err)
		return true
	}
	return false
}

// Try this three times, waiting 1s after first failure, 3s after second.
var defaultRetryBackoff = Backoff{
	Duration: 1.0 * time.Second,
	Factor:   3.0,
	Jitter:   0.1,
	Steps:    3,
}

// Useful for tests
var fastBackoff = Backoff{
	Duration: 1.0 * time.Millisecond,
	Factor:   3.0,
	Jitter:   0.1,
	Steps:    3,
}

var defaultRetryStatusCodes = []int{
	http.StatusRequestTimeout,
	http.StatusTooManyRequests, // 429: OCI distribution-spec rate limit; TooManyRequestsErrorCode is already classified temporary in transport/error.go
	http.StatusInternalServerError,
	http.StatusBadGateway,
	http.StatusServiceUnavailable,
	http.StatusGatewayTimeout,
	499, // nginx-specific, client closed request
	522, // Cloudflare-specific, connection timeout
}

const (
	defaultJobs = 4

	// ECR returns an error if n > 1000:
	// https://github.com/google/go-containerregistry/issues/1091
	defaultPageSize = 1000
)

// DefaultTransport is based on http.DefaultTransport with modifications
// documented inline below.
var DefaultTransport http.RoundTripper = &http.Transport{
	Proxy: http.ProxyFromEnvironment,
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}).DialContext,
	ForceAttemptHTTP2:     true,
	MaxIdleConns:          100,
	IdleConnTimeout:       90 * time.Second,
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
	// We usually are dealing with 2 hosts (at most), split MaxIdleConns between them.
	MaxIdleConnsPerHost: 50,
}

func makeOptions(opts ...Option) (*options, error) { _ = "STUB: not implemented"; return nil, nil }

// It is a better experience to explicitly tell a caller their auth is misconfigured
// than potentially fail silently when the correct auth is overridden by option misuse.

// transport.Wrapper is a signal that consumers are opt-ing into providing their own transport without any additional wrapping.
// This is to allow consumers full control over the transports logic, such as providing retry logic.

// Wrap the transport in something that logs requests and responses.
// It's expensive to generate the dumps, so skip it if we're writing
// to nothing.

// Using customized retry predicate if provided, and fallback to default if not.

// Wrap the transport in something that can retry network flakes.

// Wrap this last to prevent transport.New from double-wrapping.

// WithTransport is a functional option for overriding the default transport
// for remote operations.
// If transport.Wrapper is provided, this signals that the consumer does *not* want any further wrapping to occur.
// i.e. logging, retry and useragent
//
// The default transport is DefaultTransport.
func WithTransport(t http.RoundTripper) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithAuth is a functional option for overriding the default authenticator
// for remote operations.
// It is an error to use both WithAuth and WithAuthFromKeychain in the same Option set.
//
// The default authenticator is authn.Anonymous.
func WithAuth(auth authn.Authenticator) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithAuthFromKeychain is a functional option for overriding the default
// authenticator for remote operations, using an authn.Keychain to find
// credentials.
// It is an error to use both WithAuth and WithAuthFromKeychain in the same Option set.
//
// The default authenticator is authn.Anonymous.
func WithAuthFromKeychain(keys authn.Keychain) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithPlatform is a functional option for overriding the default platform
// that Image and Descriptor.Image use for resolving an index to an image.
//
// The default platform is amd64/linux.
func WithPlatform(p v1.Platform) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithContext is a functional option for setting the context in http requests
// performed by a given function. Note that this context is used for _all_
// http requests, not just the initial volley. E.g., for remote.Image, the
// context will be set on http requests generated by subsequent calls to
// RawConfigFile() and even methods on layers returned by Layers().
//
// The default context is context.Background().
func WithContext(ctx context.Context) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithJobs is a functional option for setting the parallelism of remote
// operations performed by a given function. Note that not all remote
// operations support parallelism.
//
// The default value is 4.
func WithJobs(jobs int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithUserAgent adds the given string to the User-Agent header for any HTTP
// requests. This header will also include "go-containerregistry/${version}".
//
// If you want to completely overwrite the User-Agent header, use WithTransport.
func WithUserAgent(ua string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithNondistributable includes non-distributable (foreign) layers
// when writing images, see:
// https://github.com/opencontainers/image-spec/blob/master/layer.md#non-distributable-layers
//
// The default behaviour is to skip these layers
func WithNondistributable(o *options) error { _ = "STUB: not implemented"; return nil }

// WithProgress takes a channel that will receive progress updates as bytes are written.
//
// Sending updates to an unbuffered channel will block writes, so callers
// should provide a buffered channel to avoid potential deadlocks.
func WithProgress(updates chan<- v1.Update) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithPageSize sets the given size as the value of parameter 'n' in the request.
//
// To omit the `n` parameter entirely, use WithPageSize(0).
// The default value is 1000.
func WithPageSize(size int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithRetryBackoff sets the httpBackoff for retry HTTP operations.
func WithRetryBackoff(backoff Backoff) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithRetryPredicate sets the predicate for retry HTTP operations.
func WithRetryPredicate(predicate retry.Predicate) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithRetryStatusCodes sets which http response codes will be retried.
func WithRetryStatusCodes(codes ...int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithFilter sets the filter querystring for HTTP operations.
func WithFilter(key string, value string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Reuse takes a Puller or Pusher and reuses it for remote interactions
// rather than starting from a clean slate. For example, it will reuse token exchanges
// when possible and avoid sending redundant HEAD requests.
//
// Reuse will take precedence over other options passed to most remote functions because
// most options deal with setting up auth and transports, which Reuse intetionally skips.
func Reuse[I *Puller | *Pusher](i I) Option { _ = "STUB: not implemented"; return *new(Option) }
