// Copyright 2022 The searKing Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slices_test

import (
	"fmt"
	"strconv"
	"testing"

	slices_ "github.com/searKing/golang/go/exp/slices"
	"golang.org/x/exp/slices"
)

func TestMap(t *testing.T) {
	tests := []struct {
		data []int
		want []string
	}{
		{nil, nil},
		{[]int{}, []string{}},
		{[]int{0}, []string{"0"}},
		{[]int{1, 0}, []string{"1", "0"}},
		{[]int{1, 2}, []string{"1", "2"}},
		{[]int{0, 1, 2}, []string{"0", "1", "2"}},
		{[]int{0, 1, 0, 2}, []string{"0", "1", "0", "2"}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.data), func(t *testing.T) {
			{
				got := slices_.Map(tt.data)

				if slices.Compare(got, tt.want) != 0 {
					t.Errorf("slices_.Filter(%v) = %v, want %v", tt.data, got, tt.want)
				}
			}
		})
	}
}

func TestMapFunc(t *testing.T) {
	tests := []struct {
		data []int
		want []string
	}{
		{nil, nil},
		{[]int{}, []string{}},
		{[]int{0}, []string{"0"}},
		{[]int{1, 0}, []string{"1", "0"}},
		{[]int{1, 2}, []string{"1", "2"}},
		{[]int{0, 1, 2}, []string{"0", "1", "2"}},
		{[]int{0, 1, 0, 2}, []string{"0", "1", "0", "2"}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.data), func(t *testing.T) {
			{
				got := slices_.MapFunc(tt.data, func(e int) string {
					return strconv.Itoa(e)
				})

				if slices.Compare(got, tt.want) != 0 {
					t.Errorf("slices_.FilterFunc(%v, func(e int) bool {return e != 0}) = %v, want %v", tt.data, got, tt.want)
				}
			}
		})
	}
}
