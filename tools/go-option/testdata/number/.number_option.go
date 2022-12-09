// Code generated by "go-option -type Number"; DO NOT EDIT.
// Install go-option by "go get install github.com/searKing/golang/tools/go-option"

package main

// A NumberOption sets options.
type NumberOption[T comparable] interface {
	apply(*Number[T])
}

// EmptyNumberOption does not alter the configuration. It can be embedded
// in another structure to build custom options.
//
// This API is EXPERIMENTAL.
type EmptyNumberOption[T comparable] struct{}

func (EmptyNumberOption[T]) apply(*Number[T]) {}

// NumberOptionFunc wraps a function that modifies Number[T] into an
// implementation of the NumberOption[T comparable] interface.
type NumberOptionFunc[T comparable] func(*Number[T])

func (f NumberOptionFunc[T]) apply(do *Number[T]) {
	f(do)
}

// ApplyOptions call apply() for all options one by one
func (o *Number[T]) ApplyOptions(options ...NumberOption[T]) *Number[T] {
	for _, opt := range options {
		if opt == nil {
			continue
		}
		opt.apply(o)
	}
	return o
}

// WithNumberArrayType sets arrayType in Number[T].
func WithNumberArrayType[T comparable](v [5]T) NumberOption[T] {
	return NumberOptionFunc[T](func(o *Number[T]) {
		o.arrayType = v
	})
}

// WithNumberInterfaceType sets interfaceType in Number[T].
func WithNumberInterfaceType[T comparable](v interface{}) NumberOption[T] {
	return NumberOptionFunc[T](func(o *Number[T]) {
		o.interfaceType = v
	})
}

// WithNumberMapType appends mapType in Number[T].
func WithNumberMapType[T comparable](m map[string]int64) NumberOption[T] {
	return NumberOptionFunc[T](func(o *Number[T]) {
		if o.mapType == nil {
			o.mapType = m
			return
		}
		for k, v := range m {
			o.mapType[k] = v
		}
	})
}

// WithNumberMapTypeReplace sets mapType in Number[T].
func WithNumberMapTypeReplace[T comparable](v map[string]int64) NumberOption[T] {
	return NumberOptionFunc[T](func(o *Number[T]) {
		o.mapType = v
	})
}

// WithNumberSliceType appends sliceType in Number[T].
func WithNumberSliceType[T comparable](v ...int64) NumberOption[T] {
	return NumberOptionFunc[T](func(o *Number[T]) {
		o.sliceType = append(o.sliceType, v...)
	})
}

// WithNumberSliceTypeReplace sets sliceType in Number[T].
func WithNumberSliceTypeReplace[T comparable](v ...int64) NumberOption[T] {
	return NumberOptionFunc[T](func(o *Number[T]) {
		o.sliceType = v
	})
}

// WithNumberName sets name in Number[T].
func WithNumberName[T comparable](v string) NumberOption[T] {
	return NumberOptionFunc[T](func(o *Number[T]) {
		o.name = v
	})
}
