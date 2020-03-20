// Copyright (c) 2020 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package backup

import (
	"os"
	"testing"

	. "github.com/bborbe/assert"
	"github.com/golang/glog"
)

func TestMain(m *testing.M) {
	exit := m.Run()
	glog.Flush()
	os.Exit(exit)
}

func TestRunCommand(t *testing.T) {
	err := runCommand("ls", "/")
	if err := AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
}
