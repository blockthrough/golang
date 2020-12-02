// Copyright 2020 The searKing Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package must

// Must panics if err != nil
// Deprecated: Use errors.Must instead.
func Must(err error) {
	if err == nil {
		return
	}
	panic(err)
}
