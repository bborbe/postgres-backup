// Copyright (c) 2026 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main_test

import (
	"testing"

	"github.com/onsi/gomega/gexec"
)

// TestMainCompiles verifies the main package links — without this, a build
// failure in main.go would not be caught by make test.
func TestMainCompiles(t *testing.T) {
	defer gexec.CleanupBuildArtifacts()
	if _, err := gexec.Build(".", "-mod=vendor", "-buildvcs=false"); err != nil {
		t.Fatalf("main package failed to build: %v", err)
	}
}
