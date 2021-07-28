// Copyright 2021 The searKing Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package manual

import (
	"context"

	"github.com/searKing/golang/go/net/resolver"
)

// NewBuilderWithScheme creates a new test resolver builder with the given scheme.
func NewBuilderWithScheme(scheme string) *Resolver {
	r := &Resolver{
		ResolveNowCallback: func(ctx context.Context, opts ...resolver.ResolveNowOption) {},
		scheme:             scheme,
	}
	r.ResolveAddrCallback = func(ctx context.Context, opts ...resolver.ResolveAddrOption) ([]resolver.Address, error) { return r.Addresses, nil }

	return r
}

// Resolver is also a resolver builder.
// It's build() function always returns itself.
type Resolver struct {
	ResolveAddrCallback func(ctx context.Context, opts ...resolver.ResolveAddrOption) ([]resolver.Address, error)
	// ResolveNowCallback is called when the ResolveNow method is called on the
	// resolver.  Must not be nil.  Must not be changed after the resolver may
	// be built.
	ResolveNowCallback func(ctx context.Context, opts ...resolver.ResolveNowOption)
	scheme             string

	// Addresses is the latest set of resolved addresses for the target.
	Addresses []resolver.Address

	// Fields actually belong to the resolver.
	CC             resolver.ClientConn
	bootstrapState *resolver.State
}

// InitialState adds initial state to the resolver so that UpdateState doesn't
// need to be explicitly called after Dial.
func (r *Resolver) InitialState(s resolver.State) {
	r.bootstrapState = &s
}

// Build returns itself for Resolver, because it's both a builder and a resolver.
func (r *Resolver) Build(target resolver.Target, cc resolver.ClientConn, opts ...resolver.ResolveNowOption) (resolver.Resolver, error) {
	r.CC = cc
	if r.bootstrapState != nil {
		r.UpdateState(*r.bootstrapState)
	}
	return r, nil
}

// Scheme returns the test scheme.
func (r *Resolver) Scheme() string {
	return r.scheme
}

// ResolveAddr is a noop for Resolver.
func (r *Resolver) ResolveAddr(ctx context.Context, opts ...resolver.ResolveAddrOption) ([]resolver.Address, error) {
	return r.Addresses, nil
}

// ResolveNow is a noop for Resolver.
func (r *Resolver) ResolveNow(ctx context.Context, opts ...resolver.ResolveNowOption) {
	r.ResolveNowCallback(ctx, opts...)
}

// Close is a noop for Resolver.
func (*Resolver) Close() {}

// UpdateState calls CC.UpdateState.
func (r *Resolver) UpdateState(s resolver.State) {
	_ = r.CC.UpdateState(s)
}
