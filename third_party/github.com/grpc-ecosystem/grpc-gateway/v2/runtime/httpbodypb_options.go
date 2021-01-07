// Copyright 2020 The searKing Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by "go-option -type=HTTPBodyPb"; DO NOT EDIT.

package runtime

var _default_HTTPBodyPb_value = func() (val HTTPBodyPb) { return }()

// A HTTPBodyPbOption sets options.
type HTTPBodyPbOption interface {
	apply(*HTTPBodyPb)
}

// EmptyHTTPBodyPbOption does not alter the configuration. It can be embedded
// in another structure to build custom options.
//
// This API is EXPERIMENTAL.
type EmptyHTTPBodyPbOption struct{}

func (EmptyHTTPBodyPbOption) apply(*HTTPBodyPb) {}

// HTTPBodyPbOptionFunc wraps a function that modifies HTTPBodyPb into an
// implementation of the HTTPBodyPbOption interface.
type HTTPBodyPbOptionFunc func(*HTTPBodyPb)

func (f HTTPBodyPbOptionFunc) apply(do *HTTPBodyPb) {
	f(do)
}

func (o *HTTPBodyPb) ApplyOptions(options ...HTTPBodyPbOption) *HTTPBodyPb {
	for _, opt := range options {
		if opt == nil {
			continue
		}
		opt.apply(o)
	}
	return o
}