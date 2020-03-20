// Copyright (c) 2020 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package model

type TargetDirectory string

func (b TargetDirectory) String() string {
	return string(b)
}
